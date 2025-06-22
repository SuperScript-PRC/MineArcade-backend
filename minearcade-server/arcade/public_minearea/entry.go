package public_minearea

import (
	packets_general "MineArcade-backend/minearcade-server/arcade/general"
	"MineArcade-backend/minearcade-server/clients"
	"MineArcade-backend/minearcade-server/defines"
	"MineArcade-backend/minearcade-server/defines/kick_msg"
	packet_arcade "MineArcade-backend/minearcade-server/protocol/packets/arcade"
	"fmt"
	"log"
	"log/slog"
	"sync"
)

var mmap *MineAreaMap
var players = map[string]*MineAreaPlayer{}
var player_lock = sync.Mutex{}

const PLAYER_SPAWNPOINT_X = MAP_BORDER_X / 2
const PLAYER_SPAWNPOINT_Y = MAP_BORDER_Y - CHUNK_SIZE*8 - 12

func Launch() {
	var err error
	mmap, err = ReadMapFile()
	if err != nil {
		panic(err)
	}
}

func PlayerEntry(cli *clients.ArcadeClient) {
	if !packets_general.ConfirmStartGame(cli, defines.GAMETYPE_PUBLIC_MINEAREA) {
		log.Printf("客户端 %v 未确认 StartGame", cli.AuthInfo.AccountName)
		return
	}
	if mmap == nil {
		cli.Kick("This arcade game wasn't starting.")
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
		p, err := cli.NextPacket()
		if err != nil {
			cli.Kick(kick_msg.BROKEN_PACKET)
			return
		}
		nowtime := float64time()
		switch pk := p.(type) {
		case *packet_arcade.PublicMineareaPlayerActorData:
			player.UpdateFromPacket(pk)
			if nowtime-player_move_broadcast_cd > 0.05 {
				// 避免过快收到移动数据包
				// todo: 应该直接踢出发送过快的客户端
				player_move_broadcast_cd = nowtime
				ForOtherPlayers(cli.AuthInfo.UIDStr, func(p *MineAreaPlayer) {
					p.Client.WritePacket(pk)
				})
			}
			if nowtime-player_update_chunk_cd > 0.5 {
				// 避免过于频繁地处理视野问题
				// 0.5s 更新一次区块视野
				player_update_chunk_cd = nowtime
				player.UpdatePlayerSightChunks()
			}
		case *packet_arcade.PublicMineareaBlockEvent:
			// todo: warning: can modify block without server valid checking
			err = mmap.ModifyBlock(pk.BlockX, pk.BlockY, pk.NewBlock)
			if err != nil {
				slog.Error(fmt.Sprintf(cli.AuthInfo.AccountName, "发送了不合法的方块操作:", err))
				cli.Kick(kick_msg.INVALID_PACKET)
				return
			}
			ForOtherPlayers(cli.AuthInfo.UIDStr, func(p *MineAreaPlayer) {
				p.TryUpdateBlock(pk)
			})
		case *packet_arcade.ArcadeExitGame:
			return
		default:
			cli.Kick(kick_msg.INVALID_PACKET)
			return
		}
	}

}

func AddPlayer(player *MineAreaPlayer) {
	player_lock.Lock()
	players[player.Client.AuthInfo.UIDStr] = player
	player_lock.Unlock()
	ForOtherPlayers(player.Client.AuthInfo.UIDStr, func(p *MineAreaPlayer) {
		p.Client.WritePacket(&packet_arcade.PublicMineareaPlayerActorData{
			UIDStr: player.Client.AuthInfo.UIDStr,
			X:      player.X,
			Y:      player.Y,
			Action: packet_arcade.MineAreaPlayerActionAddPlayer,
		})
	})
}

func RemovePlayer(player *MineAreaPlayer) {
	player_lock.Lock()
	delete(players, player.Client.AuthInfo.UIDStr)
	player_lock.Unlock()
	ForOtherPlayers(player.Client.AuthInfo.UIDStr, func(p *MineAreaPlayer) {
		p.Client.WritePacket(&packet_arcade.PublicMineareaPlayerActorData{
			UIDStr: player.Client.AuthInfo.UIDStr,
			X:      player.X,
			Y:      player.Y,
			Action: packet_arcade.MineAreaPlayerActionRemovePlayer,
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
		slog.Info("公共矿区地图已保存")
		SaveMapFile(mmap)
	}
}
