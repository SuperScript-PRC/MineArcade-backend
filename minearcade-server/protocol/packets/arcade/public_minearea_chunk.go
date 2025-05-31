package packets_arcade

import (
	"MineArcade-backend/minearcade-server/protocol"
	packet_define "MineArcade-backend/minearcade-server/protocol/packets/define"
)

type PublicMineAreaChunk struct {
	ChunkX int32
	ChunkY int32
	// 一个区块包含 16x16 个方块
	// []xy
	ChunkData []byte
}

func (p *PublicMineAreaChunk) ID() uint32 {
	return packet_define.IDPublicMineareaChunk
}

func (p *PublicMineAreaChunk) NetType() int8 {
	return packet_define.UDPPacket
}

func (p *PublicMineAreaChunk) Marshal(w *protocol.Writer) {
	w.Int32(p.ChunkX)
	w.Int32(p.ChunkY)
	w.Bytes(p.ChunkData)
}
