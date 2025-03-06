package clients

import (
	"MineArcade-backend/protocol"
	"MineArcade-backend/protocol/decoder"
	"MineArcade-backend/protocol/packets"
	"net"
)

type NetClient struct {
	Conn net.Conn
}

func (c *NetClient) WritePacket(p packets.ServerPacket) error {
	w := protocol.Writer{}
	p.Marshal(&w)
	_, err := c.Conn.Write(w.GetFullBytes())
	return err
}

func (c *NetClient) ReadPacket() (packets.ClientPacket, error) {
	var buf []byte
	_, err := c.Conn.Read(buf)
	if err != nil {
		return nil, err
	}
	r := protocol.Reader{}
	r.SetFullBytes(buf)
	return decoder.DecodeClientPacket(&r)
}
