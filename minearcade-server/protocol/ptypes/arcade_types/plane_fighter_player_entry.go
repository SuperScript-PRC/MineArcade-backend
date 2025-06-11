package arcade_types

import "MineArcade-backend/minearcade-server/protocol"

type PlaneFighterPlayerEntry struct {
	NickName  string
	UID       string
	RuntimeID int32
}

func (entry *PlaneFighterPlayerEntry) Marshal(w *protocol.Writer) {
	w.StringUTF(entry.NickName)
	w.StringUTF(entry.UID)
	w.Int32(entry.RuntimeID)
}
