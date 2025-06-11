package packets_arcade

import (
	"MineArcade-backend/minearcade-server/protocol"
	packet_define "MineArcade-backend/minearcade-server/protocol/packets/define"
	"MineArcade-backend/minearcade-server/protocol/ptypes/arcade_types"
)

type PlaneFighterAddActor struct {
	Actors []arcade_types.PlaneFighterActor
}

func (p *PlaneFighterAddActor) ID() uint32 {
	return packet_define.IDPlaneFighterAddActor
}

func (p *PlaneFighterAddActor) NetType() int8 {
	return packet_define.UDPPacket
}

func (p *PlaneFighterAddActor) Marshal(w *protocol.Writer) {
	protocol.WriteSlice(w, p.Actors)
}
