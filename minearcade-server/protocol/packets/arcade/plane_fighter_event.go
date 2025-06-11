package packets_arcade

import (
	"MineArcade-backend/minearcade-server/protocol"
	packet_define "MineArcade-backend/minearcade-server/protocol/packets/define"
	"MineArcade-backend/minearcade-server/protocol/ptypes/arcade_types"
)

type PlaneFighterActorEvent struct {
	Events []arcade_types.PlaneFighterActorEvent
}

func (p *PlaneFighterActorEvent) ID() uint32 {
	return packet_define.IDPlaneFighterActorEvent
}

func (p *PlaneFighterActorEvent) NetType() int8 {
	return packet_define.UDPPacket
}

func (p *PlaneFighterActorEvent) Marshal(w *protocol.Writer) {
	protocol.WriteSlice(w, p.Events)
}
