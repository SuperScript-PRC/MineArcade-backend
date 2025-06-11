package plane_fighter

import (
	"MineArcade-backend/minearcade-server/arcade/plane_fighter/define"
	"math/rand"
	"time"
)

var ran = rand.New(rand.NewSource(time.Now().UnixNano()))

func UpdateRand() {
	ran = rand.New(rand.NewSource(time.Now().UnixNano()))
}

func (stage *PlaneFighterStage) trySpawnEntities() {
	var spawnTime int32
	switch stage.Phase {
	case define.PHASE_1:
		spawnTime = define.STAGE_1_SPAWN_TICKS
	case define.PHASE_2:
		spawnTime = define.STAGE_2_SPAWN_TICKS
	}
	if stage.Phase != 0 && stage.Ticks%spawnTime == 0 {
		stage.spawnEntities()
	}
}

func (stage *PlaneFighterStage) spawnEntities() {
	println("Spawn")
	freeSpaces := stage.EntitySlotFreeSpace()
	if freeSpaces < 5 {
		return
	}
	// random entity
	rNum := ran.Int31() % 255
	switch {
	case rNum < define.EnemyPlaneWeight:
		stage.AddEntity(NewEnemyPlane(0, 0, stage.NewRuntimeID()).RandomX(ran), true)
	case rNum < define.BulletChestWeight:
		stage.AddEntity(NewBulletChest(0, 0, stage.NewRuntimeID()).RandomX(ran), true)
	case rNum < define.FIXING_PACKET_WEIGHT:
		stage.AddEntity(NewFixingPacket(0, 0, stage.NewRuntimeID()).RandomX(ran), true)
	}
}
