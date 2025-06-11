package packets_arcade

import (
	"MineArcade-backend/minearcade-server/protocol"
	packet_define "MineArcade-backend/minearcade-server/protocol/packets/define"
	"MineArcade-backend/minearcade-server/protocol/ptypes/arcade_types"
)

type PlaneFighterPlayerList struct {
	Entries []arcade_types.PlaneFighterPlayerEntry
}

func (p *PlaneFighterPlayerList) ID() uint32 {
	return packet_define.IDPlaneFighterPlayerList
}

func (p *PlaneFighterPlayerList) NetType() int8 {
	return packet_define.TCPPacket
}

func (p *PlaneFighterPlayerList) Marshal(w *protocol.Writer) {
	protocol.WriteSlice(w, p.Entries)
}
