package ptypes

import "MineArcade-backend/minearcade-server/protocol"

// 玩家排行榜数据。
type RankData struct {
	// 玩家名。
	PlayerName string
	// 玩家 UUID。
	PlayerUUID string
	// 排行榜分数。
	Score int32
	// 排行榜排名。
	Rank uint32
}

func (r *RankData) Marshal(w *protocol.Writer) {
	w.StringUTF(r.PlayerName)
	w.StringUTF(r.PlayerUUID)
	w.Int32(r.Score)
	w.UInt32(r.Rank)
}
