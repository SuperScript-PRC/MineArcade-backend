package public_minearea

import (
	"MineArcade-backend/minearcade-server/defines"
	"fmt"
	"os"
)

const MAP_CHUNK_HEIGHT = 32
const MAP_CHUNK_WIDTH = 32
const MAP_BORDER_X = MAP_CHUNK_WIDTH * CHUNK_SIZE
const MAP_BORDER_Y = MAP_CHUNK_HEIGHT * CHUNK_SIZE
const TOTAL_CHUNK_NUM = MAP_CHUNK_WIDTH * MAP_CHUNK_HEIGHT
const TOTAL_BLOCK_NUM = MAP_CHUNK_WIDTH * MAP_CHUNK_HEIGHT * CHUNK_SIZE * CHUNK_SIZE

// 地图长宽为 32x32 区块
// 按地图最左方为 x = 0
// 地图的最下方为 y = 0
type MineAreaMap struct {
	ChunkData [TOTAL_CHUNK_NUM]*Chunk
}

func NewMineAreaMap(chunkData [TOTAL_CHUNK_NUM]*Chunk) *MineAreaMap {
	return &MineAreaMap{ChunkData: chunkData}
}

func (m *MineAreaMap) InChunk(x, y int32) (*Chunk, error) {
	if x >= MAP_BORDER_X || y >= MAP_BORDER_Y {
		return nil, fmt.Errorf("x or y out of range")
	}
	chunkX := x / CHUNK_SIZE
	chunkY := y / CHUNK_SIZE
	return m.ChunkData[chunkY*MAP_CHUNK_WIDTH+chunkX], nil
}

// 较低效的 ModifyBlock.
func (m *MineAreaMap) ModifyBlock(x, y int32, blockID byte) error {
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

// 将地图数据序列化为二进制流
func (m *MineAreaMap) Marshal() [TOTAL_BLOCK_NUM]byte {
	mdata := [TOTAL_BLOCK_NUM]byte{}
	i := 0
	j := 0
	for _, chunk := range m.ChunkData {
		if chunk == nil {
			panic(fmt.Errorf("nil chunk: %v", j))
		}
		j++
		for j1 := range CHUNK_SIZE * CHUNK_SIZE {
			mdata[i] = chunk.ChunkData[j1]
			i++
		}
	}
	return mdata
}

// 从二进制流中读取地图数据
func (m *MineAreaMap) Unmarshal(mdata [TOTAL_BLOCK_NUM]byte) {
	const c = CHUNK_SIZE * CHUNK_SIZE
	for j := range MAP_CHUNK_HEIGHT * MAP_CHUNK_WIDTH {
		chunkX := j % MAP_CHUNK_HEIGHT
		chunkY := j / MAP_CHUNK_HEIGHT
		chunk := NewEmptyChunk(int32(chunkX), int32(chunkY))
		chunk.ChunkData = mdata[j*c : (j+1)*c]
		m.ModifyChunk(chunk)
	}
}

func ReadMapFile() (*MineAreaMap, error) {
	_, err := os.Stat(defines.MINEAREA_MAP_PATH)
	var file *os.File
	if os.IsNotExist(err) {
		var bs = SpawnMineAreaMap().Marshal()
		file, err = os.Create(defines.MINEAREA_MAP_PATH)
		if err != nil {
			return nil, fmt.Errorf("create map file error: %v", err)
		}
		n, err := file.Write(bs[:])
		if n == 0 {
			panic("write 0 byte info map file")
		} else if err != nil {
			return nil, fmt.Errorf("write map file error: " + err.Error())
		}
		file.Close()
	}
	file, err = os.Open(defines.MINEAREA_MAP_PATH)
	if err != nil {
		return nil, fmt.Errorf("open map file error: " + err.Error())
	}
	defer file.Close()
	content := make([]byte, TOTAL_BLOCK_NUM)
	n, err := file.Read(content)
	if err != nil {
		return nil, fmt.Errorf("read map file error: " + err.Error())
	}
	if n != TOTAL_BLOCK_NUM {
		return nil, fmt.Errorf("map file size error: %v", n)
	}
	mmap := &MineAreaMap{}
	mmap.Unmarshal([262144]byte(content))
	return mmap, nil
}

func SaveMapFile(mmap *MineAreaMap) error {
	file, err := os.OpenFile(defines.MINEAREA_MAP_PATH, os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("open map file error: " + err.Error())
	}
	defer file.Close()
	bs := mmap.Marshal()
	file.Write(bs[:])
	return nil
}
