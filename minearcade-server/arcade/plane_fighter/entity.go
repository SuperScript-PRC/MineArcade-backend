package plane_fighter

import (
	"MineArcade-backend/minearcade-server/arcade/plane_fighter/define"
	"MineArcade-backend/minearcade-server/protocol/ptypes/arcade_types"
	"math"
)

type IEntity interface {
	GetRuntimeID() int32
	GetCenterX() float32
	GetCenterY() float32
	GetHeight() int32
	GetWidth() int32
	Hurt(hp int32)
	IsDied() bool
	Remove()
	DeepRemove()
}

type BasicEntity struct {
	EntityType  int8
	RuntimeID   int32
	CenterX     float32
	CenterY     float32
	Height      int32
	Width       int32
	HP          int32
	HPMax       int32
	Removed     bool
	DeepRemoved bool
	Tick        int32
	ExtraData1  int32
	ExtraData2  int32
	ExtraData3  int32
}

func (e *BasicEntity) GetCenterX() float32 {
	return e.CenterX
}

func (e *BasicEntity) GetCenterY() float32 {
	return e.CenterY
}

func (e *BasicEntity) GetWidth() int32 {
	return e.Width
}

func (e *BasicEntity) GetHeight() int32 {
	return e.Height
}

func (e *BasicEntity) GetRuntimeID() int32 {
	return e.RuntimeID
}

func (e *BasicEntity) Cure(hp int32) (overflow int32) {
	overflow = max(0, hp-e.HPMax+e.HP)
	e.HP = min(e.HPMax, e.HP+hp)
	return
}

func (e *BasicEntity) Hurt(hp int32) {
	e.HP = max(e.HP-int32(hp), 0)
}

func (e *BasicEntity) IsDied() bool {
	return e.HP <= 0
}

// 移除实体, 且将通知客户端实体移除事件。
func (e *BasicEntity) Remove() {
	e.Removed = true
}

// 移除实体, 但是不会通知客户端实体移除事件。
func (e *BasicEntity) DeepRemove() {
	e.Removed = true
	e.DeepRemoved = true
}

func (e *BasicEntity) SimpleMarshal() arcade_types.PFStageEntity {
	return arcade_types.PFStageEntity{
		RuntimeID: e.RuntimeID,
		CenterX:   e.CenterX,
		CenterY:   e.CenterY,
	}
}

func (e *BasicEntity) InStage() bool {
	return e.CenterX >= -float32(e.Width) && e.CenterX <= float32(define.STAGE_WIDTH) &&
		e.CenterY >= -float32(e.Height) && e.CenterY <= float32(define.STAGE_HEIGHT)
}

func (e *BasicEntity) CanAttackPlayer() bool {
	return e.EntityType == define.EnemyPlane || e.EntityType == define.EnemyBullet || e.EntityType == define.EnemyLaser
}

func (e *BasicEntity) CanAttackEnemy() bool {
	return e.EntityType == define.PlayerPlane || e.EntityType == define.PlayerBullet || e.EntityType == define.PlayerLaser
}

type MovedEntity struct {
	BasicEntity
	VX float32
	VY float32
}

func (me *MovedEntity) RunTick() {
	me.Tick++
	me.CenterX += me.VX
	me.CenterY += me.VY
}

func (me *MovedEntity) Marshal() arcade_types.PlaneFighterActor {
	return arcade_types.PlaneFighterActor{
		ActorType: me.EntityType,
		RuntimeID: me.RuntimeID,
		X:         me.CenterX,
		Y:         me.CenterY,
	}
}

func (p *BasicEntity) HitTest(entity IEntity) bool {
	return math.Abs(float64(p.CenterX-entity.GetCenterX())) < float64(p.Width/2+entity.GetWidth()/2) &&
		math.Abs(float64(p.CenterY-entity.GetCenterY())) < float64(p.Height/2+entity.GetHeight()/2)
}

type Player struct {
	BasicEntity
	Score               int32
	CtrlPlayerRuntimeID int32
	Shield              int32
	Bullet              int32
	BulletMax           int32
	TotalHurt           int32
	IsFiring            bool
	Exited              chan bool
}

func (p *Player) AddBullet(bullets int32) (overflow int32) {
	overflow = max(0, bullets-(p.BulletMax-p.Bullet))
	p.Bullet = min(p.Bullet+bullets, p.BulletMax)
	return
}

func (p *Player) ReduceBullet(bullets int32) {
	p.Bullet = max(p.Bullet-bullets, 0)
}

func (p *Player) AddScore(s *PlaneFighterStage, score int32) {
	p.Score += score
	s.AddScoreEvts = append(s.AddScoreEvts, arcade_types.PlaneFighterScore{PlayerRuntimeID: p.RuntimeID, AddScore: score, TotalScore: p.Score})
}

func (p *Player) Hurt(hp int32) {
	p.BasicEntity.Hurt(hp)
	p.TotalHurt += hp
}

func (p *Player) Exit(win bool) {
	p.Exited <- win
}
