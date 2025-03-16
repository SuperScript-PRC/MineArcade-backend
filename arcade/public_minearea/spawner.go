package public_minearea

import (
	"math/rand"
	"time"
)

var rand_gen = rand.New(rand.NewSource(time.Now().UnixNano()))

func SpawnMineAreaMap() *MineAreaMap {
	t := [MAP_CHUNK_WIDTH * MAP_CHUNK_HEIGHT]*Chunk{}
	m_map := &MineAreaMap{ChunkData: t}
	// 在距离顶端 10 区块的范围内填充充满空气的区块
	for height_i := range 10 {
		for width_i := range MAP_CHUNK_WIDTH {
			m_map.ModifyChunk(NewEmptyChunk(uint(width_i), uint(MAP_CHUNK_HEIGHT-height_i-1)))
		}
	}
	// 剩下的部分直接随机矿物方块
	for height_i := range MAP_CHUNK_HEIGHT - 10 {
		for width_i := range MAP_CHUNK_WIDTH {
			chunk := NewEmptyChunk(uint(width_i), uint(height_i))
			for x := range CHUNK_SIZE {
				for y := range CHUNK_SIZE {
					chunk.ModifyBlock(uint(x), uint(y), RandomMineBlock())
				}
			}
			m_map.ModifyChunk(chunk)
		}
	}
	return m_map
}

func RandomMineBlock() byte {
	// 绿宝石: 0.5% 钻石: 0.5% 金矿石: 2% 红石: 4% 青金石: 2.2%
	// 铁矿石: 5% 煤矿石: 8%
	// 其他: 石头
	res := rand_gen.Intn(1000)
	if res < 5 {
		return EmeraldOre
	} else if res < 10 {
		return DiamondOre
	} else if res < 30 {
		return GoldOre
	} else if res < 70 {
		return RedstoneOre
	} else if res < 92 {
		return LapisOre
	} else if res < 142 {
		return IronOre
	} else if res < 200 {
		return CoalOre
	} else {
		return Stone
	}
}
