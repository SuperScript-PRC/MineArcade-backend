package packets_general

import (
	"MineArcade-backend/minearcade-server/protocol"
	packet_define "MineArcade-backend/minearcade-server/protocol/packets/define"
)

type UDPConnection struct {
	VerifyToken string
}

func (p *UDPConnection) ID() uint32 {
	return packet_define.IDUDPConnection
}

func (p *UDPConnection) NetType() int8 {
	return packet_define.UDPPacket
}

func (p *UDPConnection) Unmarshal(r *protocol.Reader) {
	r.StringUTF(&p.VerifyToken)
}
