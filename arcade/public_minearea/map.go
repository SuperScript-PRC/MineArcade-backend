package public_minearea

import (
	"fmt"
)

const MAP_CHUNK_HEIGHT = 32
const MAP_CHUNK_WIDTH = 32

const MAP_BORDER_X = MAP_CHUNK_WIDTH * CHUNK_SIZE
const MAP_BORDER_Y = MAP_CHUNK_HEIGHT * CHUNK_SIZE

const TOTAL_BLOCK_NUM = MAP_CHUNK_WIDTH * MAP_CHUNK_HEIGHT * CHUNK_SIZE * CHUNK_SIZE

// 地图长宽为 32x32 区块
// 按地图最左方为 x = 0
// 地图的最下方为 y = 0
type MineAreaMap struct {
	ChunkData [MAP_CHUNK_WIDTH * MAP_CHUNK_HEIGHT]*Chunk
}

func (m *MineAreaMap) InChunk(x, y uint) (*Chunk, error) {
	if x >= MAP_BORDER_X || y >= MAP_BORDER_Y {
		return nil, fmt.Errorf("x or y out of range")
	}
	chunkX := x / CHUNK_SIZE
	chunkY := y / CHUNK_SIZE
	return m.ChunkData[chunkY*MAP_CHUNK_WIDTH+chunkX], nil
}

// 较低效的 ModifyBlock.
func (m *MineAreaMap) ModifyBlock(x, y uint, blockID byte) error {
	chunk, err := m.InChunk(x, y)
	if err != nil {
		return err
	}
	chunk.ModifyBlock(x-chunk.ChunkX*CHUNK_SIZE, y-chunk.ChunkY*CHUNK_SIZE, blockID)
	return nil
}

func (m *MineAreaMap) ModifyChunk(chunk *Chunk) error {
	if chunk.ChunkX >= MAP_CHUNK_WIDTH || chunk.ChunkY >= MAP_CHUNK_HEIGHT {
		return fmt.Errorf("ChunkX: %v or ChunkY: %v out of range", chunk.ChunkX, chunk.ChunkY)
	}
	m.ChunkData[chunk.ChunkY*MAP_CHUNK_WIDTH+chunk.ChunkX] = chunk
	return nil
}

// 将地图数据序列化成二进制流
func (m *MineAreaMap) Marshal() [TOTAL_BLOCK_NUM]byte {
	var mdata [TOTAL_BLOCK_NUM]byte
	i := 0
	for _, chunk := range m.ChunkData {
		for j := range CHUNK_SIZE * CHUNK_SIZE {
			mdata[i] = chunk.ChunkData[j]
			i++
		}
	}
	return mdata
}

// 从二进制流中读取地图数据
func (m *MineAreaMap) Unmarshal(mdata [TOTAL_BLOCK_NUM]byte) {
	i := 0
	for _, chunk := range &m.ChunkData {
		for j := range CHUNK_SIZE * CHUNK_SIZE {
			chunk.ChunkData[j] = mdata[i]
			i++
		}
	}
}
