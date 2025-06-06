package packets_arcade

import (
	"MineArcade-backend/minearcade-server/protocol"
	packet_define "MineArcade-backend/minearcade-server/protocol/packets/define"
)

// 公共矿区的方块变动
type PublicMineareaBlockEvent struct {
	BlockX   int32
	BlockY   int32
	NewBlock byte
}

func (p *PublicMineareaBlockEvent) ID() uint32 {
	return packet_define.IDPublicMineareaBlockEvent
}

func (p *PublicMineareaBlockEvent) NetType() int8 {
	return packet_define.UDPPacket
}

func (p *PublicMineareaBlockEvent) Marshal(w *protocol.Writer) {
	w.Int32(p.BlockX)
	w.Int32(p.BlockY)
	w.UInt8(p.NewBlock)
}

func (p *PublicMineareaBlockEvent) Unmarshal(r *protocol.Reader) {
	r.Int32(&p.BlockX)
	r.Int32(&p.BlockY)
	r.UInt8(&p.NewBlock)
}
