package plane_fighter

import (
	"MineArcade-backend/minearcade-server/arcade/plane_fighter/define"
	"MineArcade-backend/minearcade-server/protocol/ptypes/arcade_types"
	"math"
	"math/rand"
)

func StageEntityToActor(e *MovedEntity) arcade_types.PlaneFighterActor {
	return arcade_types.PlaneFighterActor{
		ActorType: e.EntityType,
		RuntimeID: e.RuntimeID,
		X:         e.CenterX,
		Y:         e.CenterY,
	}
}

func (e1 *BasicEntity) Distance(e2 IEntity) float64 {
	return math.Sqrt(math.Pow(float64(e1.CenterX-e2.GetCenterX()), 2) + math.Pow(float64(e1.CenterY-e2.GetCenterY()), 2))
}

func (e1 *BasicEntity) DistanceX(e2 IEntity) float32 {
	return e1.CenterX - e2.GetCenterX()
}

func (e1 *BasicEntity) DistanceY(e2 IEntity) float32 {
	return e1.CenterY - e2.GetCenterY()
}

func (e1 *MovedEntity) RandomX(ran *rand.Rand) *MovedEntity {
	rangeLeft := e1.Width / 2
	rangeRight := define.STAGE_WIDTH - e1.Width/2
	e1.CenterX = float32(ran.Int31()%(rangeRight-rangeLeft) + rangeLeft)
	return e1
}
