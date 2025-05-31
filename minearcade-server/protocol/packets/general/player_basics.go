package packets_general

import (
	"MineArcade-backend/minearcade-server/protocol"
	packet_define "MineArcade-backend/minearcade-server/protocol/packets/define"
)

// 描述玩家基本信息。
// 在玩家登录成功, 发送了 ClientLoginResp 后发送。
type PlayerBasics struct {
	// 玩家的昵称
	Nickname string
	// 玩家的 UID
	UID string
	// 玩家的钱数
	Money float64
	// 玩家的体力值
	Power int32
	// 玩家的点数
	Points int32
	// 玩家的等级
	Level int32
	// 玩家当前的经验值
	Exp int32
	// 玩家升级所需的经验值
	ExpUpgrade int32
}

func (p *PlayerBasics) ID() uint32 {
	return packet_define.IDPlayerBasics
}

func (p *PlayerBasics) NetType() int8 {
	return packet_define.TCPPacket
}

func (p *PlayerBasics) Marshal(w *protocol.Writer) {
	w.StringUTF(p.Nickname)
	w.StringUTF(p.UID)
	w.Double(p.Money)
	w.Int32(p.Power)
	w.Int32(p.Points)
	w.Int32(p.Level)
	w.Int32(p.Exp)
	w.Int32(p.ExpUpgrade)
}
