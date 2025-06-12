package define

const (
	ENEMY_PLANE_WEIGHT   = 160
	BULLET_CHEST_WEIGHT  = 40
	FIXING_PACKET_WEIGHT = 20
)

const (
	EnemyPlaneWeight   = ENEMY_PLANE_WEIGHT
	BulletChestWeight  = EnemyPlaneWeight + BULLET_CHEST_WEIGHT
	FixingPacketWeight = BulletChestWeight + FIXING_PACKET_WEIGHT
)
