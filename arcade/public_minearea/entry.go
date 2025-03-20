package public_minearea

import (
	"MineArcade-backend/arcade/general"
	"MineArcade-backend/clients"
	"MineArcade-backend/defines"
	"MineArcade-backend/protocol/packets"
	"fmt"
	"sync"

	"github.com/pterm/pterm"
)

var mmap *MineAreaMap
var players = map[string]*MineAreaPlayer{}
var player_lock = sync.Mutex{}

const PLAYER_SPAWNPOINT_X = MAP_BORDER_X / 2
const PLAYER_SPAWNPOINT_Y = MAP_BORDER_Y - CHUNK_SIZE*10 - 12

func Launch() {
	var err error
	mmap, err = ReadMapFile()
	if err != nil {
		panic(fmt.Errorf("read map file error: %v", err))
	}
	defer func() {
		SaveMapFile(mmap)
		pterm.Success.Println("公共矿区地图已保存")
	}()
}

func PlayerEntry(cli *clients.NetClient) {
	if !general.ConfirmStartGame(cli, defines.GAMETYPE_PUBLIC_MINEAREA) {
		return
	}
	if mmap == nil {
		cli.Kick("This arcade game isn't starting.")
		return
	}
	player := NewPlayer(mmap, cli, PLAYER_SPAWNPOINT_X, PLAYER_SPAWNPOINT_Y)
	players[cli.AuthInfo.UIDStr] = player
	defer RemovePlayer(player)
	AddPlayer(player)
	player.UpdatePlayerSightChunks()
	var player_move_broadcast_cd float64
	var player_update_chunk_cd float64
	for {
		p, err := cli.ReadNextPacket()
		if err != nil {
			cli.Kick("Broken packet")
			return
		}
		nowtime := float64time()
		if pk, ok := p.(*packets.PublicMineareaPlayerActorData); ok {
			player.UpdateFromPacket(pk)
			if nowtime-player_move_broadcast_cd > 0.05 {
				// 避免过快收到移动数据包
				player_move_broadcast_cd = nowtime
				ForOtherPlayers(cli.AuthInfo.UIDStr, func(p *MineAreaPlayer) {
					p.Client.WritePacket(pk)
				})
			}
			if nowtime-player_update_chunk_cd > 0.5 {
				// 避免过于频繁地处理视野问题
				player_update_chunk_cd = nowtime
				player.UpdatePlayerSightChunks()
			}
		} else if pk, ok := p.(*packets.PublicMineareaBlockEvent); ok {
			// TODO: can modify block without server valid checking
			mmap.ModifyBlock(pk.BlockX, pk.BlockY, pk.NewBlock)
			ForOtherPlayers(cli.AuthInfo.UIDStr, func(p *MineAreaPlayer) {
				p.TryUpdateBlock(pk)
			})
		} else {
			cli.Kick("不合法的操作")
			return
		}
	}
}

func AddPlayer(player *MineAreaPlayer) {
	player_lock.Lock()
	players[player.Client.AuthInfo.UIDStr] = player
	player_lock.Unlock()
	ForOtherPlayers(player.Client.AuthInfo.UIDStr, func(p *MineAreaPlayer) {
		p.Client.WritePacket(&packets.PublicMineareaPlayerActorData{
			UIDStr: player.Client.AuthInfo.UIDStr,
			X:      player.X,
			Y:      player.Y,
			Action: packets.MineAreaPlayerActionAddPlayer,
		})
	})
}

func RemovePlayer(player *MineAreaPlayer) {
	player_lock.Lock()
	delete(players, player.Client.AuthInfo.UIDStr)
	player_lock.Unlock()
	ForOtherPlayers(player.Client.AuthInfo.UIDStr, func(p *MineAreaPlayer) {
		p.Client.WritePacket(&packets.PublicMineareaPlayerActorData{
			UIDStr: player.Client.AuthInfo.UIDStr,
			X:      player.X,
			Y:      player.Y,
			Action: packets.MineAreaPlayerActionRemovePlayer,
		})
	})
}

func ForAllPlayers(f func(*MineAreaPlayer)) {
	defer player_lock.Unlock()
	player_lock.Lock()
	for _, player := range players {
		f(player)
	}
}

func ForOtherPlayers(senderUID string, f func(*MineAreaPlayer)) {
	defer player_lock.Unlock()
	player_lock.Lock()
	for uuid, player := range players {
		if senderUID != uuid {
			f(player)
		}
	}
}

func Exit() {
	if mmap != nil {
		SaveMapFile(mmap)
	}
}
