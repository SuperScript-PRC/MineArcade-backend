package packets_arcade

import (
	"MineArcade-backend/minearcade-server/protocol"
	packet_define "MineArcade-backend/minearcade-server/protocol/packets/define"
	"MineArcade-backend/minearcade-server/protocol/ptypes/arcade_types"
)

type PlaneFighterPlayerStatuses struct {
	Statuses []arcade_types.PFPlayerStatus
}

func (p *PlaneFighterPlayerStatuses) ID() uint32 {
	return packet_define.IDPlaneFighterPlayerStatuses
}

func (p *PlaneFighterPlayerStatuses) NetType() int8 {
	return packet_define.UDPPacket
}

func (p *PlaneFighterPlayerStatuses) Marshal(w *protocol.Writer) {
	protocol.WriteSlice(w, p.Statuses)
}
