package arcade_types

import "MineArcade-backend/minearcade-server/protocol"

type PFStageEntity struct {
	RuntimeID int32
	CenterX   float32
	CenterY   float32
}

func (e *PFStageEntity) Marshal(w *protocol.Writer) {
	w.Int32(e.RuntimeID)
	w.Float32(e.CenterX)
	w.Float32(e.CenterY)
}
