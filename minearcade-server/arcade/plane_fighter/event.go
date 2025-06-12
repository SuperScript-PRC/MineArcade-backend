package plane_fighter

import (
	"MineArcade-backend/minearcade-server/protocol/ptypes/arcade_types"
)

const (
	EVENT_ADD_ENTITY = iota + 1
	EVENT_REMOVE_ENTITY
	EVENT_DIED
	EVENT_TNT_EXPLODED
	EVENT_COLORFUL_EXPLODE
)

type Event = arcade_types.PlaneFighterActorEvent
