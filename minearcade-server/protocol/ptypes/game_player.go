package ptypes

import "MineArcade-backend/minearcade-server/protocol"

// 代表一个玩家，无论处于哪种游戏中。
// 其只具有名称和 UUID 两种属性。
type GamePlayer struct {
	Name string
	UUID string
}

func (p *GamePlayer) Marshal(w *protocol.Writer) {
	w.StringUTF(p.Name)
	w.StringUTF(p.UUID)
}
