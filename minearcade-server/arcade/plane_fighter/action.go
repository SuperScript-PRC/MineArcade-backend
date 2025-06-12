package plane_fighter

import (
	"MineArcade-backend/minearcade-server/arcade/plane_fighter/define"
)

func EntityRunTick(s *PlaneFighterStage, e *MovedEntity) {
	e.RunTick()
	switch e.EntityType {
	case define.TNT:
		if TNTAddTimer(e) >= define.TNTExplodeTick {
			TNTExplode(s, e)
		}
	case define.EnemyPlane:
		if e.Tick%define.EnemyPlaneShootDelay == 0 {
			EnemyPlaneShoot(s, e)
		}
	}
}

func PlayerRunTick(s *PlaneFighterStage, p *Player) {
	if p.IsFiring && p.Bullet > 0 {
		if p.Bullet > 0 {
			p.Tick++
			if p.Tick%define.PlayerShootDelay == 0 {
				PlayerShoot(s, p)
			}
		}

	}
}

func (s *PlaneFighterStage) Weapon2EntityAction(e1 *MovedEntity, e2 *MovedEntity) {
	if !e1.CanAttackEnemy() || e1.Removed || e2.Removed {
		return
	}
	switch e2.EntityType {
	case define.EnemyPlane:
		switch e1.EntityType {
		case define.PlayerBullet:
			if e1.HitTest(e2) {
				BulletHitEnemyPlane(s, e1, e2)
			}
		case define.PlayerMissile:
			if e1.HitTest(e2) {
				MissileHitEnemyPlane(s, e1, e2)
			}
		}
	case define.BulletChest:
		if e1.HitTest(e2) {
			HitBulletChest(s, e1, e2)
		}
	case define.FixingPacket:
		if e1.HitTest(e2) {
			HitFixingPacket(s, e1, e2)
		}
	}
}

func (s *PlaneFighterStage) Player2EntityAction(p *Player, e *MovedEntity) {
	if e.Removed || !p.HitTest(e) {
		return
	}
	switch e.EntityType {
	case define.EnemyPlane:
		EnemyPlaneHitPlayer(s, p, e)
	case define.EnemyBullet:
		BulletHitPlayer(s, p, e)
	case define.EnemyLaser:
		LaserHitPlayer(s, p, e)
	case define.BulletChest:
		e.Remove()
		overflow := p.AddBullet(define.BULLET_CHEST_BULLET_COUNT)
		p.AddScore(s, overflow*define.BULLET_OVERFLOW_BONUS)
	case define.FixingPacket:
		e.Remove()
		overflow := p.Cure(define.FIXING_PACKET_CURE)
		p.AddScore(s, overflow*define.CURE_OVERFLOW_BONUS)
	}
}

func PlayerShoot(s *PlaneFighterStage, p *Player) {
	s.AddEntity(NewPlayerBullet(p.CenterX, p.CenterY+float32(p.Height)/2, s.NewRuntimeID(), p.RuntimeID), define.DEBUG_SHOW_BULLET)
	p.Bullet--
}

func EnemyPlaneShoot(s *PlaneFighterStage, e *MovedEntity) {
	//s.AddEntity(NewEnemyBullet(e.CenterX, e.CenterY+float32(e.Height)/2, s.NewRuntimeID()), define.DEBUG_SHOW_BULLET)
}

func TNTAddTimer(e *MovedEntity) int32 {
	e.ExtraData1 += 1
	return e.ExtraData1
}

func TNTExplode(s *PlaneFighterStage, e *MovedEntity) {
	e.Remove()
	for _, e2 := range s.Entities {
		dist := e.Distance(e2)
		if e != e2 && dist < 200 {
			if dist == 0 {
				dist = 1
			}
			hurt := int32(400 / dist)
			x_delta := e.DistanceX(e2)
			y_delta := e.DistanceY(e2)
			if x_delta == 0 {
				x_delta = 1
			}
			if y_delta == 0 {
				y_delta = 1
			}
			x_speed_delta := 20 / x_delta
			y_speed_delta := 20 / y_delta
			e.VX += x_speed_delta
			e.VY += y_speed_delta
			Hurt(s, e, hurt)
		}
	}
	for _, p := range s.PlayerPlanes {
		dist := p.Distance(e)
		if dist < 200 {
			if dist == 0 {
				dist = 1
			}
			hurt := int32(400 / dist)
			x_delta := e.DistanceX(p)
			y_delta := e.DistanceY(p)
			if x_delta == 0 {
				x_delta = 1
			}
			if y_delta == 0 {
				y_delta = 1
			}
			x_speed_delta := 20 / x_delta
			y_speed_delta := 20 / y_delta
			e.VX += x_speed_delta
			e.VY += y_speed_delta
			Hurt(s, e, hurt)
		}
	}
	s.AddEvent(Event{EntityRuntimeID: e.RuntimeID, EventID: EVENT_TNT_EXPLODED})
}

// P2E actions
func BulletHitEnemyPlane(s *PlaneFighterStage, bullet *MovedEntity, eplane *MovedEntity) {
	bullet.DeepRemove()
	if Hurt(s, eplane, define.PLAYER_BULLET_HURT) {
		ToScore(s, bullet, define.SCORE_PLAYER_SHOOT_ENEMY_PLANE)
	}
}

func HitBulletChest(s *PlaneFighterStage, bullet *MovedEntity, bchest *MovedEntity) {
	bullet.DeepRemove()
	bchest.DeepRemove()
	ToScore(s, bullet, define.SCORE_PLAYER_SHOOT_BULLET_CHEST)
	p := s.GetPlayer(bullet.ExtraData1)
	if p != nil {
		overflow := p.AddBullet(define.BULLET_CHEST_BULLET_COUNT)
		p.AddScore(s, overflow*define.BULLET_OVERFLOW_BONUS)
	}
	s.Events = append(s.Events, Event{EventID: EVENT_COLORFUL_EXPLODE, EntityRuntimeID: bchest.RuntimeID})
}

func HitFixingPacket(s *PlaneFighterStage, bullet *MovedEntity, fpacket *MovedEntity) {
	bullet.DeepRemove()
	fpacket.DeepRemove()
	ToScore(s, bullet, define.SCORE_PLAYER_SHOOT_FIXING_PACKET)
	p := s.GetPlayer(bullet.ExtraData1)
	if p != nil {
		overflow := p.Cure(define.FIXING_PACKET_CURE)
		p.AddScore(s, overflow*define.CURE_OVERFLOW_BONUS)
	}
	s.Events = append(s.Events, Event{EventID: EVENT_COLORFUL_EXPLODE, EntityRuntimeID: fpacket.RuntimeID})
}

func LaserHitEnemyPlane(s *PlaneFighterStage, laser *MovedEntity, eplane *MovedEntity) {
	if Hurt(s, eplane, define.PLAYER_LASER_HURT) {
		ToScore(s, laser, define.SCORE_PLAYER_SHOOT_ENEMY_PLANE)
	}
}

func MissileHitEnemyPlane(s *PlaneFighterStage, missile *MovedEntity, eplane *MovedEntity) {
	missile.Remove()
	if Hurt(s, eplane, define.PLAYER_MISSILE_HURT) {
		ToScore(s, missile, define.SCORE_PLAYER_SHOOT_ENEMY_PLANE)
	}
}

// E2P actions

func EnemyPlaneHitPlayer(s *PlaneFighterStage, p *Player, e *MovedEntity) {
	s.Events = append(s.Events, Event{EventID: EVENT_DIED, EntityRuntimeID: e.RuntimeID})
	e.DeepRemove()
	Hurt(s, p, define.ENEMY_PLANE_HURT)
}

func BulletHitPlayer(s *PlaneFighterStage, p *Player, e *MovedEntity) {
	e.Remove()
	Hurt(s, p, define.ENEMY_BULLET_HURT)
}

func LaserHitPlayer(s *PlaneFighterStage, p *Player, e *MovedEntity) {
	e.Remove()
	Hurt(s, p, define.ENEMY_LASER_HURT)
}

// Simple actions

func Hurt(s *PlaneFighterStage, e IEntity, damage int32) (isDied bool) {
	if isDied = e.IsDied(); isDied {
		e.Remove()
		return
	}
	e.Hurt(damage)
	if isDied = e.IsDied(); isDied {
		s.Events = append(s.Events, Event{EntityRuntimeID: e.GetRuntimeID(), EventID: EVENT_DIED})
		e.DeepRemove()
	}
	return
}

func ToScore(s *PlaneFighterStage, killerEntity *MovedEntity, score int32) {
	if killerEntity.ExtraData1 != 0 {
		player := s.GetPlayer(killerEntity.ExtraData1)
		if player != nil {
			player.AddScore(s, define.SCORE_PLAYER_SHOOT_ENEMY_PLANE)
		}
	}
}

// Player actions

func (s *PlaneFighterStage) PlayerStartFire(playerRuntimeID int32) {
	if p := s.GetPlayer(playerRuntimeID); p != nil {
		p.IsFiring = true
	}
}

func (s *PlaneFighterStage) PlayerStopFire(playerRuntimeID int32) {
	if p := s.GetPlayer(playerRuntimeID); p != nil {
		p.IsFiring = false
	}
}

func (s *PlaneFighterStage) PlayerSyncPosition(playerRuntimeID int32, x float32, y float32) {
	if p := s.GetPlayer(playerRuntimeID); p != nil {
		p.CenterX = x
		p.CenterY = y
	}
}
