package packets

import (
	"MineArcade-backend/protocol"
)

type PublicMineAreaChunk struct {
	ChunkX int32
	ChunkY int32
	// 一个区块包含 64x64 个方块
	// []xy
	ChunkData []byte
}

func (p *PublicMineAreaChunk) ID() uint32 {
	return IDPublicMineAreaChunk
}

func (p *PublicMineAreaChunk) Marshal(w *protocol.Writer) {
	w.Int32(p.ChunkX)
	w.Int32(p.ChunkY)
	w.Bytes(p.ChunkData)
}
