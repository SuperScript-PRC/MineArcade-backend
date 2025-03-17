package public_minearea

import (
	"MineArcade-backend/clients"
	"MineArcade-backend/protocol/packets"
	"math"
)

const PLAYER_SIGHT = 64

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
		cx, cy := GetChunkXYByIndex(uint(i))
		if math.Abs(player.X-float64(cx+HALF_CHUNK_SIZE)) < PLAYER_SIGHT && math.Abs(player.Y-float64(cy+HALF_CHUNK_SIZE)) < PLAYER_SIGHT {
			player.loadChunk(cx, cy)
		} else {
			player.unloadChunk(cx, cy)
		}
	}
}

func (player *MineAreaPlayer) Teleport(x, y float64) {
	player.X = x
	player.Y = y
	player.Client.WritePacket(&packets.PublicMineareaPlayerActorData{
		UUIDStr: player.Client.AuthInfo.UUIDStr,
		X:       x,
		Y:       y,
		Action:  0,
	})
}

func (player *MineAreaPlayer) loadChunk(x, y uint) {
	index := y*MAP_CHUNK_WIDTH + x
	if !player.VisiChunks[index] {
		pk := player.Map.ChunkData[index].ConvertToPacket()
		player.Client.WritePacket(&pk)
		player.VisiChunks[index] = true
	}
}

func (player *MineAreaPlayer) unloadChunk(chunkX, chunkY uint) {
	player.VisiChunks[chunkY*MAP_CHUNK_WIDTH+chunkX] = false
}

func (player *MineAreaPlayer) ChunkLoaded(chunkX, chunkY uint) bool {
	return player.VisiChunks[chunkY*MAP_CHUNK_WIDTH+chunkX]
}

func (player *MineAreaPlayer) TryUpdateBlock(pk *packets.PublicMineareaBlockEvent) {
	chunk_x, chunk_y := ConvertToChunkXY(uint(pk.BlockX), uint(pk.BlockY))
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
