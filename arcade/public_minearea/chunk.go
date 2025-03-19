package public_minearea

import (
	"MineArcade-backend/protocol/packets"
	"fmt"
)

const CHUNK_SIZE = 16
const HALF_CHUNK_SIZE = CHUNK_SIZE / 2

// 一个区块包含 16x16 个方块
type Chunk struct {
	// 区块 X 坐标。使用相对坐标为以区块为单位的坐标系。
	ChunkX uint
	// 区块 Y 坐标。使用相对坐标为以区块为单位的坐标系。
	ChunkY uint
	// 区块信息。为 []yx 存储
	ChunkData []byte
}

func (c *Chunk) ModifyBlock(relative_x, relative_y uint, blockID byte) error {
	if relative_x > CHUNK_SIZE || relative_y > CHUNK_SIZE {
		return fmt.Errorf("relative_x or relative_y out of range %v", CHUNK_SIZE)
	}
	c.ChunkData[relative_y*CHUNK_SIZE+relative_x] = blockID
	return nil
}

func (c *Chunk) ConvertToPacket() packets.PublicMineAreaChunk {
	return packets.PublicMineAreaChunk{ChunkX: int32(c.ChunkX), ChunkY: int32(c.ChunkY), ChunkData: c.ChunkData}
}

func (c *Chunk) CenterXY() (uint, uint) {
	return c.ChunkX*CHUNK_SIZE + HALF_CHUNK_SIZE, c.ChunkY*CHUNK_SIZE + HALF_CHUNK_SIZE
}

func NewChunk(chunkX, chunkY uint, chunkData []byte) *Chunk {
	return &Chunk{ChunkX: chunkX, ChunkY: chunkY, ChunkData: chunkData}
}

func NewEmptyChunk(chunkX, chunkY uint) *Chunk {
	bts := [CHUNK_SIZE * CHUNK_SIZE]byte{}
	return NewChunk(chunkX, chunkY, bts[:])
}

func AlignToChunk(x, y uint) (uint, uint) {
	return x / CHUNK_SIZE * CHUNK_SIZE, y / CHUNK_SIZE * CHUNK_SIZE
}

func ConvertToChunkXY(x, y uint) (uint, uint) {
	return x / CHUNK_SIZE, y / CHUNK_SIZE
}

func GetChunkXYByIndex(index uint) (uint, uint) {
	chunkX := index % MAP_CHUNK_WIDTH
	chunkY := index / MAP_CHUNK_WIDTH
	return chunkX, chunkY
}

func GetChunkCenterXYByIndex(index int) (int, int) {
	chunkX := index % MAP_CHUNK_WIDTH
	chunkY := index / MAP_CHUNK_WIDTH
	return chunkX*CHUNK_SIZE + HALF_CHUNK_SIZE, chunkY*CHUNK_SIZE + HALF_CHUNK_SIZE
}
