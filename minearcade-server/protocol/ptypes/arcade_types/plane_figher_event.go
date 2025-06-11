package arcade_types

import "MineArcade-backend/minearcade-server/protocol"

type PlaneFighterActorEvent struct {
	EventID         int8
	EntityRuntimeID int32
}

func (e *PlaneFighterActorEvent) Marshal(w *protocol.Writer) {
	w.Int8(e.EventID)
	w.Int32(e.EntityRuntimeID)
}
