package public_minearea

import (
	"MineArcade-backend/clients"
	"MineArcade-backend/protocol/packets"
	"fmt"

	"github.com/pterm/pterm"
)

var mmap *MineAreaMap
var players map[string]*MineAreaPlayer

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
	if mmap == nil {
		cli.Kick("This arcade game isn't starting.")
		return
	}
	player := NewPlayer(mmap, cli, PLAYER_SPAWN_X, PLAYER_SPAWN_Y)
	players[cli.AuthInfo.UUIDStr] = player
	defer func() {
		delete(players, cli.AuthInfo.UUIDStr)
	}()
	var player_move_broadcast_cd float32
	var player_update_chunk_cd float32
	for {
		p, err := cli.ReadNextPacket()
		if err != nil {
			cli.Kick("Broken packet")
			return
		}
		nowtime := float32time()
		if pk, ok := p.(*packets.PublicMineareaPlayerActorData); ok {
			if nowtime-player_move_broadcast_cd > 0.05 {
				// 避免过快收到移动数据包
				player_move_broadcast_cd = nowtime
				ForOtherPlayers(cli.AuthInfo.UUIDStr, func(p *MineAreaPlayer) {
					p.Client.WritePacket(pk)
				})
			}
			if nowtime-player_update_chunk_cd > 0.5 {
				// 避免过于频繁地处理视野问题
				player.UpdatePlayerSightChunks()
			}
		} else if pk, ok := p.(*packets.PublicMineareaBlockEvent); ok {
			// TODO: can modify block without server valid checking
			mmap.ModifyBlock(uint(pk.BlockX), uint(pk.BlockY), pk.NewBlock)
			ForOtherPlayers(cli.AuthInfo.UUIDStr, func(p *MineAreaPlayer) {
				p.TryUpdateBlock(pk)
			})
		} else {

		}
	}
}

func ForAllPlayers(f func(*MineAreaPlayer)) {
	for _, player := range players {
		f(player)
	}
}

func ForOtherPlayers(senderUUID string, f func(*MineAreaPlayer)) {
	for uuid, player := range players {
		if senderUUID != uuid {
			f(player)
		}
	}
}
