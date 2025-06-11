package packets_arcade

import (
	"MineArcade-backend/minearcade-server/protocol"
	packet_define "MineArcade-backend/minearcade-server/protocol/packets/define"
)

const (
	PFPEventStartFire = iota + 1
	PFPEventStopFire
)

type PlaneFighterPlayerEvent struct {
	EventID int8
	Value   int32
}

func (p *PlaneFighterPlayerEvent) ID() uint32 {
	return packet_define.IDPlaneFighterPlayerEvent
}

func (p *PlaneFighterPlayerEvent) NetType() int8 {
	return packet_define.UDPPacket
}

func (p *PlaneFighterPlayerEvent) Unmarshal(r *protocol.Reader) {
	r.Int8(&p.EventID)
	r.Int32(&p.Value)
}
