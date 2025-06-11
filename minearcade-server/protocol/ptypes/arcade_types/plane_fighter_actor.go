package arcade_types

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
	ActorType int8
	RuntimeID int32
	X         float32
	Y         float32
}

// Simply marshal for stage.
func (a *PlaneFighterActor) Marshal(w *protocol.Writer) {
	w.Int8(a.ActorType)
	w.Int32(a.RuntimeID)
	w.Double(float64(a.X))
	w.Double(float64(a.Y))
}
