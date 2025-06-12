package arcade_types

import "MineArcade-backend/minearcade-server/protocol"

type PFPlayerStatus struct {
	RuntimeID int32
	HP        int32
	Bullets   int32
}

func (status *PFPlayerStatus) Marshal(w *protocol.Writer) {
	w.Int32(status.RuntimeID)
	w.Int32(status.HP)
	w.Int32(status.Bullets)
}
