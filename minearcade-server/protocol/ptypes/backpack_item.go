package ptypes

import "MineArcade-backend/minearcade-server/protocol"

type Item struct {
	// 物品的唯一 ID
	ID int32
	// 物品数量
	Amount int32
}

func (it *Item) Marshal(w *protocol.Writer) {
	w.Int32(it.ID)
	w.Int32(it.Amount)
}
