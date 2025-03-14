package packets

import "MineArcade-backend/protocol"

// 公共矿区的方块变动。
type PublicMineAreaBlockEvent struct {
	BlockX   int32
	BlockY   int32
	Action   byte
	NewBlock byte
}

func (p *PublicMineAreaBlockEvent) ID() uint32 {
	return IDPublicMineAreaBlockEvent
}

func (p *PublicMineAreaBlockEvent) Marshal(w *protocol.Writer) {
	w.Int32(p.BlockX)
	w.Int32(p.BlockY)
	w.UInt8(p.Action)
	w.UInt8(p.NewBlock)
}

func (p *PublicMineAreaBlockEvent) Unmarshal(r *protocol.Reader) {
	r.Int32(&p.BlockX)
	r.Int32(&p.BlockY)
	r.UInt8(&p.Action)
	r.UInt8(&p.NewBlock)
}
