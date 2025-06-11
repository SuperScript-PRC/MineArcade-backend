package packets_arcade

import (
	"MineArcade-backend/minearcade-server/protocol"
	packet_define "MineArcade-backend/minearcade-server/protocol/packets/define"
	"MineArcade-backend/minearcade-server/protocol/ptypes/arcade_types"
)

type PlaneFighterScores struct {
	Scores []arcade_types.PlaneFighterScore
}

func (p *PlaneFighterScores) ID() uint32 {
	return packet_define.IDPlaneFighterScores
}

func (p *PlaneFighterScores) NetType() int8 {
	return packet_define.UDPPacket
}

func (p *PlaneFighterScores) Marshal(w *protocol.Writer) {
	protocol.WriteSlice(w, p.Scores)
}
