package arcade_types

import "MineArcade-backend/minearcade-server/protocol"

type ArcadeGameScoreDetail struct {
	ScoreID int8
	Score   int32
}

func (p *ArcadeGameScoreDetail) Marshal(w *protocol.Writer) {
	w.Int8(p.ScoreID)
	w.Int32(p.Score)
}
