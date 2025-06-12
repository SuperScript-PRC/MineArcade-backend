package define

const (
	PlayerPlane   = iota // 玩家
	EnemyPlane           // 小型敌机
	PlayerBullet         // 玩家子弹
	PlayerLaser          // 玩家镭射
	PlayerMissile        // 玩家导弹
	TNT                  // TNT
	EnemyBullet          // 敌方子弹
	EnemyLaser           // 敌方镭射
	BulletChest          // 弹匣箱
	FixingPacket         // 修补包
	BossPlane
)

func DontUpload(entity_type int32) bool {
	return entity_type == PlayerBullet || entity_type == EnemyBullet
}

func IsSyncEntity(entity_type int8) bool {
	return entity_type != PlayerBullet && entity_type != EnemyBullet
}
