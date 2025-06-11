package packets_arcade

import (
	"MineArcade-backend/minearcade-server/protocol"
	packet_define "MineArcade-backend/minearcade-server/protocol/packets/define"
)

type PlaneFighterPlayerMove struct {
	X float32
	Y float32
}

func (p *PlaneFighterPlayerMove) ID() uint32 {
	return packet_define.IDPlaneFighterPlayerMove
}

func (p *PlaneFighterPlayerMove) NetType() int8 {
	return packet_define.UDPPacket
}

func (p *PlaneFighterPlayerMove) Unmarshal(r *protocol.Reader) {
	r.Float32(&p.X)
	r.Float32(&p.Y)
}
