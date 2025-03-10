package clients

import (
	"MineArcade-backend/protocol"
	"MineArcade-backend/protocol/decoder"
	"MineArcade-backend/protocol/packets"
	"fmt"
	"net"
)

type NetClient struct {
	Conn     net.Conn
	IPString string
	Username string
	r        protocol.Reader
}

func (c *NetClient) WritePacket(p packets.ServerPacket) error {
	w := protocol.Writer{}
	w.Int32(int32(p.ID()))
	p.Marshal(&w)
	_, err := c.Conn.Write(w.GetFullBytes())
	return err
}

func (c *NetClient) ReadPackets() ([]packets.ClientPacket, error) {
	buf := make([]byte, 1048576)
	var pks []packets.ClientPacket
	n, err := c.Conn.Read(buf)
	if err != nil {
		return nil, err
	}
	r := protocol.Reader{}
	r.AddFullBytes(buf, n)
	for {
		if r.End() {
			break
		}
		pk, err := decoder.DecodeClientPacket(&r)
		if err != nil {
			return nil, err
		}
		pks = append(pks, pk)
	}
	return pks, nil
}

func (c *NetClient) ReadNextPacket() (packets.ClientPacket, error) {
	if !c.r.End() {
		pk, err := decoder.DecodeClientPacket(&c.r)
		if err != nil {
			return nil, err
		}
		return pk, nil
	} else {
		buf := make([]byte, 524288) // 512KB
		n, err := c.Conn.Read(buf)
		if err != nil {
			return nil, err
		}
		if n == 0 {
			return nil, fmt.Errorf("EOF")
		}
		c.r.SetFullBytes(buf, n)
		return c.ReadNextPacket()
	}
}
