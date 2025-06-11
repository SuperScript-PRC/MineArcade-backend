package packets_arcade

import (
	"MineArcade-backend/minearcade-server/protocol"
	packet_define "MineArcade-backend/minearcade-server/protocol/packets/define"
)

type PlaneFighterTimer struct {
	SecondsLeft int32
}

func (p *PlaneFighterTimer) ID() uint32 {
	return packet_define.IDPlaneFighterTimer
}

func (p *PlaneFighterTimer) NetType() int8 {
	return packet_define.UDPPacket
}

func (p *PlaneFighterTimer) Marshal(w *protocol.Writer) {
	w.Int32(p.SecondsLeft)
}
