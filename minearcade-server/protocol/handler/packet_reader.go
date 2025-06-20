package handler

import (
	"MineArcade-backend/minearcade-server/configs"
	"MineArcade-backend/minearcade-server/protocol"
	"MineArcade-backend/minearcade-server/protocol/decoder"
	"MineArcade-backend/minearcade-server/protocol/packets"
	packet_define "MineArcade-backend/minearcade-server/protocol/packets/define"
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"time"
)

type PacketListener func(packets.ClientPacket)

type PacketReader struct {
	TCPConn     net.Conn
	UDPConn     *net.UDPConn
	UDPConnAddr *net.UDPAddr
	rawPkQueue  chan []byte
	pkQueue     chan packets.ClientPacket
	errQueue    chan error
	pkListeners [][]PacketListener
}

func NewPacketReader(tcp_conn net.Conn, udp_conn *net.UDPConn, udp_addr *net.UDPAddr) *PacketReader {
	r := &PacketReader{
		TCPConn:     tcp_conn,
		UDPConn:     udp_conn,
		UDPConnAddr: udp_addr,
		rawPkQueue:  make(chan []byte, 256),
		pkQueue:     make(chan packets.ClientPacket, configs.CLIENT_PACKET_BUFSIZE),
		errQueue:    make(chan error, 5),
		pkListeners: make([][]PacketListener, packet_define.MaxPacketID),
	}
	return r
}

// 激活 PacketReader。
func (pr *PacketReader) Active() {
	go pr.acceptTCPPacketsFromClient()
}

// 开始接受数据包, 并将其堆积到 pkQueue 中。
func (pr *PacketReader) acceptTCPPacketsFromClient() {
	// defer func() {
	// 	pr.pkQueue <- nil
	// }()
	for {
		var packetSize int32
		err := binary.Read(pr.TCPConn, binary.BigEndian, &packetSize)
		if err != nil {
			pr.errQueue <- err
			return
		}
		bs := make([]byte, packetSize)
		n, err := pr.TCPConn.Read(bs)
		// todo: ?
		if err != nil {
			pr.errQueue <- err
			return
		}
		if int(packetSize) != n {
			pr.errQueue <- fmt.Errorf("packet size error: need %d bytes, got %d", packetSize, n)
			return
		}
		reader := protocol.NewReader(bs)
		pk, err := decoder.DecodeClientPacket(&reader)
		if err != nil {
			pr.errQueue <- err
			return
		}
		var packet_is_listened = false
		// 优先监听常监听包
		// todo: 可能造成阻塞
		for _, pk_listener := range pr.pkListeners[pk.ID()] {
			packet_is_listened = true
			pk_listener(pk)
		}
		if !packet_is_listened {
			pr.pkQueue <- pk
		}
	}
}

func (pr *PacketReader) ReceiveUDPBytePacket(pkBytes []byte) {
	var packetSize int32
	buf := bytes.NewBuffer(pkBytes)
	err := binary.Read(buf, binary.BigEndian, &packetSize)
	if err != nil {
		pr.errQueue <- err
		return
	}
	pkBytes = buf.Bytes()
	if len(pkBytes) != int(packetSize) {
		pr.errQueue <- fmt.Errorf("packet size error: need %d bytes, got %d", packetSize, len(pkBytes))
	}
	reader := protocol.NewReader(pkBytes)
	pk, err := decoder.DecodeClientPacket(&reader)
	if err != nil {
		pr.errQueue <- err
		return
	}
	var packet_is_listened = false
	for _, pk_listener := range pr.pkListeners[pk.ID()] {
		packet_is_listened = true
		pk_listener(pk)
	}
	if !packet_is_listened {
		pr.pkQueue <- pk
	}
}

func (pr *PacketReader) NextPacket() (packets.ClientPacket, error) {
	select {
	case err := <-pr.errQueue:
		return nil, err
	case pk := <-pr.pkQueue:
		return pk, nil
	}
}

func (pr *PacketReader) NextPacketWithInterrupt(c chan bool) (packets.ClientPacket, error, bool) {
	select {
	case err := <-pr.errQueue:
		return nil, err, false
	case pk := <-pr.pkQueue:
		return pk, nil, false
	case <-c:
		return nil, nil, true
	}
}

func (pr *PacketReader) WaitForPacket(pkID int, timeout time.Duration) (pk packets.ClientPacket, getted bool) {
	ch := make(chan packets.ClientPacket)
	receiver := func(pk packets.ClientPacket) {
		ch <- pk
	}
	pr.AddPacketListener(pkID, receiver)
	defer pr.RemovePacketListener(pkID, receiver)
	for {
		select {
		case pk := <-ch:
			return pk, true
		case <-time.After(timeout):
			return nil, false
		}
	}
}

func (pr *PacketReader) AddPacketListener(id int, listener PacketListener) {
	pr.pkListeners[id] = append(pr.pkListeners[id], listener)
}

func (pr *PacketReader) RemovePacketListener(id int, listener PacketListener) {
	for i, l := range pr.pkListeners[id] {
		if &l == &listener {
			pr.pkListeners[id] = append(pr.pkListeners[id][:i], pr.pkListeners[id][i+1:]...)
			break
		}
	}
}
