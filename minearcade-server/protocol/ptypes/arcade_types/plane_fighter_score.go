package arcade_types

import "MineArcade-backend/minearcade-server/protocol"

type PlaneFighterScore struct {
	PlayerRuntimeID int32
	AddScore        int32
	TotalScore      int32
}

func (p *PlaneFighterScore) Marshal(w *protocol.Writer) {
	w.Int32(p.PlayerRuntimeID)
	w.Int32(p.AddScore)
	w.Int32(p.TotalScore)
}
