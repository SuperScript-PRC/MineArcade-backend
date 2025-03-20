package public_minearea

import (
	"MineArcade-backend/clients"
	"MineArcade-backend/protocol/packets"
	"math"
)

const PLAYER_SIGHT_X = 48
const PLAYER_SIGHT_Y = 32

type MineAreaPlayer struct {
	// X, Y: 在以一方块为单位的坐标系中的位置
	Map        *MineAreaMap
	Client     *clients.NetClient
	X          float64
	Y          float64
	VisiChunks []bool
}

func (player *MineAreaPlayer) UpdateFromPacket(p *packets.PublicMineareaPlayerActorData) {
	player.X = p.X
	player.Y = p.Y
}

func (player *MineAreaPlayer) UpdatePlayerSightChunks() {
	for i := range MAP_CHUNK_HEIGHT * MAP_CHUNK_WIDTH {
		chunkX, chunkY := GetChunkXYByIndex(int32(i))
		if math.Abs(player.X-float64(chunkX*CHUNK_SIZE+HALF_CHUNK_SIZE)) < PLAYER_SIGHT_X && math.Abs(player.Y-float64(chunkY*CHUNK_SIZE+HALF_CHUNK_SIZE)) < PLAYER_SIGHT_Y {
			player.loadChunk(chunkX, chunkY)
		} else {
			player.unloadChunk(chunkX, chunkY)
		}
	}
}

func (player *MineAreaPlayer) Teleport(x, y float64) {
	player.X = x
	player.Y = y
	player.Client.WritePacket(&packets.PublicMineareaPlayerActorData{
		UIDStr: player.Client.AuthInfo.UIDStr,
		X:      x,
		Y:      y,
		Action: 0,
	})
}

func (player *MineAreaPlayer) loadChunk(chunkX, chunkY int32) {
	index := chunkY*MAP_CHUNK_WIDTH + chunkX
	// special
	if !player.VisiChunks[index] {
		pk := player.Map.ChunkData[index].ConvertToPacket()
		player.Client.WritePacket(&pk)
		player.VisiChunks[index] = true
	}
}

func (player *MineAreaPlayer) unloadChunk(chunkX, chunkY int32) {
	index := chunkY*MAP_CHUNK_WIDTH + chunkX
	if player.VisiChunks[index] {
		player.Client.WritePacket(&packets.PublicMineAreaChunk{
			ChunkX:    chunkX,
			ChunkY:    chunkY,
			ChunkData: []byte{},
		})
	}
	player.VisiChunks[index] = false
}

func (player *MineAreaPlayer) ChunkLoaded(chunkX, chunkY int32) bool {
	return player.VisiChunks[chunkY*MAP_CHUNK_WIDTH+chunkX]
}

func (player *MineAreaPlayer) TryUpdateBlock(pk *packets.PublicMineareaBlockEvent) {
	chunk_x, chunk_y := ConvertToChunkXY((pk.BlockX), (pk.BlockY))
	if player.ChunkLoaded(chunk_x, chunk_y) {
		player.Client.WritePacket(pk)
	}
}

func NewPlayer(mmap *MineAreaMap, cli *clients.NetClient, x, y float64) *MineAreaPlayer {
	player := &MineAreaPlayer{
		Map:        mmap,
		Client:     cli,
		VisiChunks: make([]bool, MAP_CHUNK_HEIGHT*MAP_CHUNK_WIDTH),
		X:          x,
		Y:          y,
	}
	player.Teleport(x, y)
	return player
}
