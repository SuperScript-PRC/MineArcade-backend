package plane_fighter

import (
	"MineArcade-backend/minearcade-server/arcade/plane_fighter/define"
	"MineArcade-backend/minearcade-server/protocol/ptypes/arcade_types"
)

// 雷霆战机的舞台描述。
// 舞台用于管理和存放雷霆战机游戏中的实体 (包括玩家),
// 并管理游戏刻。

type PlaneFighterStage struct {
	Entities         [define.STAGE_MAX_ENTITIES]*MovedEntity
	PlayerPlanes     []*Player
	Events           []Event
	NewActors        []MovedEntity
	AddScoreEvts     []arcade_types.PlaneFighterScore
	RuntimeIDCounter int32
	Ticks            int32
	TicksLeft        int32
	Phase            int32
	Exited           chan bool
}

func NewStage() *PlaneFighterStage {
	return &PlaneFighterStage{
		Entities:         [define.STAGE_MAX_ENTITIES]*MovedEntity{},
		PlayerPlanes:     []*Player{},
		RuntimeIDCounter: (-1 << 31) + 1,
		Events:           []Event{},
		NewActors:        []MovedEntity{},
		AddScoreEvts:     []arcade_types.PlaneFighterScore{},
		Ticks:            0,
		TicksLeft:        define.STAGE_PREPARE_REMAIN_TIME_TICKS,
		Phase:            0,
		Exited:           make(chan bool, 4),
	}
}

func (s *PlaneFighterStage) RunTick() {
	s.trySpawnEntities()
	s.addTick()
	if s.isEnded() {
		s.Exit()
		return
	}
	for _, p := range s.PlayerPlanes {
		PlayerRunTick(s, p)
		if p.IsDied() {
			p.Exit(false)
		}
	}
	for _, e := range s.Entities {
		if e == nil {
			continue
		}
		EntityRunTick(s, e)
		for _, p := range s.PlayerPlanes {
			if p.HitTest(e) {
				s.Player2EntityAction(p, e)
			}
		}
		for _, e2 := range s.Entities {
			if e2 != nil && e != e2 && !e.Removed && !e2.Removed {
				s.Weapon2EntityAction(e, e2)
			}
		}
	}
	s.entitiesGC()
}

func (s *PlaneFighterStage) AddPlayer(player *Player) {
	s.PlayerPlanes = append(s.PlayerPlanes, player)
}

func (s *PlaneFighterStage) EntitySlotFreeSpace() int {
	var freeSpaceLeft int
	for _, e := range s.Entities {
		if e == nil {
			freeSpaceLeft++
		}
	}
	return freeSpaceLeft
}

func (s *PlaneFighterStage) NewRuntimeID() int32 {
	s.RuntimeIDCounter++
	return s.RuntimeIDCounter
}

func (s *PlaneFighterStage) AddEntity(e *MovedEntity, makeEvent bool) bool {
	for i, e2 := range s.Entities {
		if e2 == nil {
			s.Entities[i] = e
			if makeEvent {
				s.NewActors = append(s.NewActors, *e)
			}
			return true
		}
	}
	return false
}

func (s *PlaneFighterStage) RemoveEntity(e IEntity, makeEvent bool) {
	for i, e2 := range s.Entities {
		if e == e2 {
			s.Entities[i] = nil
			if makeEvent {
				s.AddEvent(Event{EventID: EVENT_REMOVE_ENTITY, EntityRuntimeID: e.GetRuntimeID()})
			}
			return
		}
	}
}

func (s *PlaneFighterStage) AddEvent(e Event) {
	s.Events = append(s.Events, e)
}

func (s *PlaneFighterStage) GetPlayer(runtimeID int32) *Player {
	for _, p := range s.PlayerPlanes {
		if p.RuntimeID == runtimeID {
			return p
		}
	}
	return nil
}

func (s *PlaneFighterStage) WaitExited() {
	<-s.Exited
}

func (s *PlaneFighterStage) entitiesGC() {
	for i, e := range s.Entities {
		if e != nil && (e.Removed || !e.InStage()) {
			s.Entities[i] = nil
			if define.IsSyncEntity(e.EntityType) && !e.DeepRemoved {
				s.AddEvent(Event{EventID: EVENT_REMOVE_ENTITY, EntityRuntimeID: e.GetRuntimeID()})
			}
		}
	}
}

// 正常退出舞台
func (s *PlaneFighterStage) Exit() {
	for _, entity := range s.Entities {
		if entity == nil {
			continue
		}
		s.Events = append(s.Events, Event{EventID: EVENT_COLORFUL_EXPLODE, EntityRuntimeID: entity.GetRuntimeID()})
		entity.DeepRemove()
	}
	s.Exited <- true
	s.entitiesGC()
}
