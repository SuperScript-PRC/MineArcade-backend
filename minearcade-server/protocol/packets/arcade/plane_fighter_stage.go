package packets_arcade

import (
	"MineArcade-backend/minearcade-server/protocol"
	packet_define "MineArcade-backend/minearcade-server/protocol/packets/define"
	"MineArcade-backend/minearcade-server/protocol/ptypes/arcade_types"
)

type PlaneFighterStage struct {
	Players  []arcade_types.PFStageEntity
	Entities []arcade_types.PFStageEntity
}

func (p *PlaneFighterStage) ID() uint32 {
	return packet_define.IDPlaneFighterStage
}

func (p *PlaneFighterStage) NetType() int8 {
	return packet_define.UDPPacket
}

func (p *PlaneFighterStage) Marshal(w *protocol.Writer) {
	protocol.WriteSlice(w, p.Players)
	protocol.WriteSlice(w, p.Entities)
}
