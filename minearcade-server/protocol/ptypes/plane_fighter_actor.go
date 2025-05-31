package ptypes

import "MineArcade-backend/minearcade-server/protocol"

const (
	PFActorTypeEnemy = iota
	PFActorTypeBomb
	PFActorTypeFuel
	PFActorTypeBullet
	PFActorTypeGold
	PFActorTypeHP
)

type PlaneFighterActor struct {
	ActorType  int32
	RuntimeID  int32
	X          float32
	Y          float32
	ActorEvent int32
}

// Simply marshal for stage.
func (a *PlaneFighterActor) Marshal(w *protocol.Writer) {
	w.Int32(a.ActorType)
	w.Int32(a.RuntimeID)
	w.Double(float64(a.X))
	w.Double(float64(a.Y))
	w.Int32(a.ActorEvent)
}
