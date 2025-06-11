package plane_fighter

import "MineArcade-backend/minearcade-server/arcade/plane_fighter/define"

// 各类实体初始化。

func NewPlayer(runtimeid int32) *Player {
	return &Player{
		BasicEntity: BasicEntity{
			EntityType: define.PlayerPlane,
			RuntimeID:  runtimeid,
			CenterX:    define.STAGE_WIDTH / 2,
			CenterY:    define.STAGE_HEIGHT - define.PLAYER_HEIGHT,
			Height:     define.PLAYER_HEIGHT,
			Width:      define.PLAYER_WIDTH,
			HP:         define.PLAYER_HP,
		},
		CtrlPlayerRuntimeID: runtimeid,
		Shield:              0,
		Bullet:              50,
		BulletMax:           100,
		IsFiring:            false,
	}
}

func NewPlayerBullet(x float32, y float32, runtimeid int32, pusherRtID int32) *MovedEntity {
	return &MovedEntity{
		BasicEntity: BasicEntity{
			EntityType: define.PlayerBullet,
			RuntimeID:  runtimeid,
			CenterX:    x,
			CenterY:    y,
			Height:     define.PLAYER_BULLET_HEIGHT,
			Width:      define.PLAYER_BULLET_WIDTH,
			HP:         define.PLAYER_BULLET_HP,
			HPMax:      define.PLAYER_BULLET_HP,
			ExtraData1: pusherRtID,
		},
		VX: define.PLAYER_BULLET_INIT_VX,
		VY: define.PLAYER_BULLET_INIT_VY,
	}
}

func NewEnemyBullet(x float32, y float32, runtimeid int32) *MovedEntity {
	return &MovedEntity{
		BasicEntity: BasicEntity{
			EntityType: define.EnemyBullet,
			RuntimeID:  runtimeid,
			CenterX:    x,
			CenterY:    y,
			Height:     define.ENEMY_BULLET_HEIGHT,
			Width:      define.ENEMY_BULLET_WIDTH,
			HP:         define.ENEMY_BULLET_HP,
			HPMax:      define.ENEMY_BULLET_HP,
		},
		VX: define.ENEMY_BULLET_INIT_VX,
		VY: define.ENEMY_BULLET_INIT_VY,
	}
}

//func NewPlayerMissile()

func NewPlayerLaser(x float32, y float32, runtimeid int32, pusherRtID int32) *MovedEntity {
	return &MovedEntity{
		BasicEntity: BasicEntity{
			EntityType: define.PlayerLaser,
			RuntimeID:  runtimeid,
			CenterX:    x,
			CenterY:    y,
			Height:     define.PLAYER_LASER_HEIGHT,
			Width:      define.PLAYER_LASER_WIDTH,
			HP:         define.PLAYER_LASER_HP,
			HPMax:      define.PLAYER_LASER_HP,
			ExtraData1: pusherRtID,
		},
		VX: define.PLAYER_LASER_INIT_VX,
		VY: define.PLAYER_LASER_INIT_VY,
	}
}

func NewEnemyLaser(x float32, y float32, runtimeid int32) *MovedEntity {
	return &MovedEntity{
		BasicEntity: BasicEntity{
			EntityType: define.EnemyLaser,
			RuntimeID:  runtimeid,
			CenterX:    x,
			CenterY:    y,
			Height:     define.ENEMY_LASER_HEIGHT,
			Width:      define.ENEMY_LASER_WIDTH,
			HP:         define.ENEMY_LASER_HP,
			HPMax:      define.ENEMY_LASER_HP,
		},
		VX: define.ENEMY_LASER_INIT_VX,
		VY: define.ENEMY_LASER_INIT_VY,
	}
}

func NewEnemyPlane(x float32, y float32, runtimeid int32) *MovedEntity {
	return &MovedEntity{
		BasicEntity: BasicEntity{
			EntityType: define.EnemyPlane,
			RuntimeID:  runtimeid,
			CenterX:    x,
			CenterY:    y,
			Height:     define.ENEMY_PLANE_HEIGHT,
			Width:      define.ENEMY_PLANE_WIDTH,
			HP:         define.ENEMY_PLANE_HP,
			HPMax:      define.ENEMY_PLANE_HP,
			Removed:    false,
		},
		VX: define.ENEMY_PLANE_INIT_VX,
		VY: define.ENEMY_PLANE_INIT_VY,
	}
}

func NewBulletChest(x float32, y float32, runtimeid int32) *MovedEntity {
	return &MovedEntity{
		BasicEntity: BasicEntity{
			EntityType: define.BulletChest,
			RuntimeID:  runtimeid,
			CenterX:    x,
			CenterY:    y,
			Height:     define.BULLET_CHEST_HEIGHT,
			Width:      define.BULLET_CHEST_WIDTH,
			HP:         1,
			HPMax:      1,
			Removed:    false,
		},
		VX: define.ENEMY_PLANE_INIT_VX,
		VY: define.ENEMY_PLANE_INIT_VY,
	}
}

func NewFixingPacket(x float32, y float32, runtimeid int32) *MovedEntity {
	return &MovedEntity{
		BasicEntity: BasicEntity{
			EntityType: define.FixingPacket,
			RuntimeID:  runtimeid,
			CenterX:    x,
			CenterY:    y,
			Height:     define.FIXING_PACKET_HEIGHT,
			Width:      define.FIXING_PACKET_WIDTH,
			HP:         1,
			HPMax:      1,
			Removed:    false,
		},
		VX: define.ENEMY_PLANE_INIT_VX,
		VY: define.ENEMY_PLANE_INIT_VY,
	}
}
