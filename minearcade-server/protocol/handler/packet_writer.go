package handler

import (
	"MineArcade-backend/minearcade-server/configs"
	"MineArcade-backend/minearcade-server/protocol"
	"MineArcade-backend/minearcade-server/protocol/packets"
	packet_define "MineArcade-backend/minearcade-server/protocol/packets/define"
	"bytes"
	"encoding/binary"
	"net"
)

type PacketWriter struct {
	TCPConn     net.Conn
	UDPConn     *net.UDPConn
	UDPConnAddr *net.UDPAddr
	pkQueue     chan packets.ServerPacket
	err         error
}

func NewPacketWriter(tcp_conn net.Conn, udp_conn *net.UDPConn, udp_addr *net.UDPAddr) *PacketWriter {
	return &PacketWriter{
		TCPConn:     tcp_conn,
		UDPConn:     udp_conn,
		UDPConnAddr: udp_addr,
		pkQueue:     make(chan packets.ServerPacket, configs.SERVER_PACKET_BUFSIZE),
		err:         nil,
	}
}

// 激活 PacketWriter。
func (pw *PacketWriter) Active() {
	go pw.sendPacketsToClient()
}

// 将缓存的待发出的数据包发送给客户端。
func (pw *PacketWriter) sendPacketsToClient() {
	for {
		pk := <-pw.pkQueue
		w := protocol.NewWriter()
		w_final := bytes.NewBuffer([]byte{})
		w.Int32(int32(pk.ID()))
		pk.Marshal(&w)
		err := binary.Write(w_final, binary.BigEndian, int32(w.Size()))
		if err != nil {
			pw.err = err
			return
		}
		w_final.Write(w.GetFullBytes())
		switch pk.NetType() {
		case packet_define.TCPPacket:
			_, err = pw.TCPConn.Write(w_final.Bytes())
			if err != nil {
				pw.err = err
				return
			}
		case packet_define.UDPPacket:
			_, err := pw.UDPConn.WriteToUDP(w_final.Bytes(), pw.UDPConnAddr)
			if err != nil {
				pw.err = err
				return
			}
		}
	}
}

func (pw *PacketWriter) WritePacket(pk packets.ServerPacket) error {
	if pw.err != nil {
		return pw.err
	}
	pw.pkQueue <- pk
	return nil
}
