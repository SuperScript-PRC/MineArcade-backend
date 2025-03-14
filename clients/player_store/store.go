package player_store

type PlayerStore struct {
	Points     int32
	Power      int32
	Level      int32
	Exp        int32
	ExpUpgrade int32
	Money      float64
	Bagpack    []Item
}

type Item struct {
	// 物品的唯一 ID
	ID int32
	// 物品数量
	Amount int32
	// 附加值
	ItemData string
}
