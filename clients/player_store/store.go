package player_store

import (
	"MineArcade-backend/protocol"
	"fmt"
	"sync"
)

type PlayerStore struct {
	bp         sync.Mutex
	Nickname   string
	Points     int32
	Power      int32
	Level      int32
	Exp        int32
	ExpUpgrade int32
	Money      float64
	Backpack   []Item
}

type Item struct {
	// 物品的唯一 ID
	ID int32
	// 物品数量
	Amount int32
	// 附加值
	ItemData string
}

func (it *Item) Marshal(w *protocol.Writer) {
	w.Int32(it.ID)
	w.Int32(it.Amount)
	w.StringUTF(it.ItemData)
}

func (it *Item) Unmarshal(r *protocol.Reader) {
	r.Int32(&it.ID)
	r.Int32(&it.Amount)
	r.StringUTF(&it.ItemData)
}

func (ps *PlayerStore) Marshal(w *protocol.Writer) {
	w.StringUTF(ps.Nickname)
	w.Int32(ps.Points)
	w.Int32(ps.Power)
	w.Int32(ps.Level)
	w.Int32(ps.Exp)
	w.Int32(ps.ExpUpgrade)
	w.Double(ps.Money)
	protocol.WriteSlice(w, ps.Backpack)
}

func (ps *PlayerStore) Unmarshal(r *protocol.Reader) {
	r.StringUTF(&ps.Nickname)
	r.Int32(&ps.Points)
	r.Int32(&ps.Power)
	r.Int32(&ps.Level)
	r.Int32(&ps.Exp)
	r.Int32(&ps.ExpUpgrade)
	r.Double(&ps.Money)
	protocol.ReadSlice(r, &ps.Backpack)
}

func (it *Item) Add(amount int32) {
	it.Amount += amount
}

func (it *Item) Reduce(amount int32) bool {
	it.Amount -= amount
	if it.Amount < 0 {
		panic(fmt.Errorf("Item: %v count = %v, < 0", it.ID, it.Amount))
	}
	return it.Amount == 0
}

func (ps *PlayerStore) AddItem(itemID int32, amount int32) (new bool) {
	defer ps.bp.Unlock()
	ps.bp.Lock()
	new = false
	for i, item := range ps.Backpack {
		if item.ID == itemID {
			item.Add(amount)
			ps.Backpack[i] = item
			return
		}
	}
	new = true
	ps.Backpack = append(ps.Backpack, Item{ID: itemID, Amount: amount})
	return
}

func (ps *PlayerStore) ReduceItem(itemID int32, amount int32) (none bool) {
	defer ps.bp.Unlock()
	ps.bp.Lock()
	none = false
	for i, item := range ps.Backpack {
		if item.ID == itemID {
			none = item.Reduce(amount)
			if none {
				ps.Backpack = append(ps.Backpack[:i], ps.Backpack[i+1:]...)
			}
			return
		}
	}
	panic(fmt.Errorf("item id %v not in backpack", itemID))
}

func NewPlayerStore() *PlayerStore {
	return &PlayerStore{
		Points:     0,
		Power:      0,
		Level:      0,
		Exp:        0,
		ExpUpgrade: 0,
		Money:      0,
		Backpack:   []Item{},
	}
}
