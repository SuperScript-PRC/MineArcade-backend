package define

const (
	ENEMY_PLANE_WEIGHT   = 5000
	BULLET_CHEST_WEIGHT  = 6
	FIXING_PACKET_WEIGHT = 2
)

const (
	EnemyPlaneWeight   = ENEMY_PLANE_WEIGHT
	BulletChestWeight  = EnemyPlaneWeight + BULLET_CHEST_WEIGHT
	FixingPacketWeight = BulletChestWeight + FIXING_PACKET_WEIGHT
)
