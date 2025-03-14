package public_minearea

import (
	"MineArcade-backend/clients"
	"MineArcade-backend/protocol/packets"
	"math"
)

const PLAYER_SIGHT = 64

type MineAreaPlayer struct {
	Map     MineAreaMap
	Client  clients.NetClient
	UUIDStr string
	Nicknme string
	// X, Y: 在以一方块为单位的坐标系中的位置
	X          float64
	Y          float64
	VisiChunks []bool
}

func (player *MineAreaPlayer) UpdateFromPacket(p *packets.PlayerActorData) {
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

func (player *MineAreaPlayer) loadChunk(x, y uint) {
	index := y*MAP_CHUNK_WIDTH + x
	if !player.VisiChunks[index] {
		pk := player.Map.ChunkData[index].ConvertToPacket()
		player.Client.WritePacket(&pk)
		player.VisiChunks[index] = true
	}
}

func (player *MineAreaPlayer) unloadChunk(x, y uint) {
	player.VisiChunks[y*MAP_CHUNK_WIDTH+x] = false
}
