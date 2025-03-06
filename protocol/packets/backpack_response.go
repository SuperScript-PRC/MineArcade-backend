package packets

import "MineArcade-backend/protocol"

// 响应客户端的查询背包请求。
type BackpackResponse struct {
	// 背包物品列表
	Items []Item
}

type Item struct {
	// 物品的唯一 ID
	ID int32
	// 物品数量
	Amount int32
}

func (p *BackpackResponse) ID() uint32 {
	return IDBackpackResponse
}

func MarshalBackpackItem(w *protocol.Writer, it *Item) {
	w.Int32(it.ID)
	w.Int32(it.Amount)
}

func (p *BackpackResponse) Marshal(w *protocol.Writer) {
	protocol.WriteSliceWithNewMarshaler(w, p.Items, MarshalBackpackItem)
}
