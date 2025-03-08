package clients

import (
	"MineArcade-backend/protocol"
	"MineArcade-backend/protocol/decoder"
	"MineArcade-backend/protocol/packets"
	"net"
)

type NetClient struct {
	Conn     net.Conn
	Username string
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
	r.SetFullBytes(buf, n)
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
