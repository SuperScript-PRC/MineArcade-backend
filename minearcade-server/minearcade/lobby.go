package minearcade

import (
	"MineArcade-backend/minearcade-server/clients"
	kickmsg "MineArcade-backend/minearcade-server/defines/kick_msg"
	packets_arcade "MineArcade-backend/minearcade-server/protocol/packets/arcade"
	"fmt"
	"log/slog"
)

func MainLobbyEntry(cli *clients.ArcadeClient) {
	for {
		if !cli.Online {
			slog.Info(fmt.Sprintf("玩家 %s 已下线", cli.FormatNameWithUUID()))
			return
		}
		slog.Info(fmt.Sprintf("等待玩家 %s 加入游戏", cli.FormatNameWithUUID()))
		p, err := cli.NextPacket()
		if err != nil {
			slog.Error(fmt.Sprintf("玩家 %s 连接断开, 错误: %v", cli.FormatNameWithUUID(), err))
			cli.Kick(kickmsg.BROKEN_PACKET)
			return
		}
		if pk, ok := p.(*packets_arcade.ArcadeEntryRequest); ok {
			GoArcade(cli, pk.ArcadeGameType, pk.RequestUUID)
		} else {
			slog.Error(fmt.Sprintf("客户端发送不正确数据包: %d", p.ID()))
			cli.Kick(kickmsg.INVALID_PACKET)
		}
	}
}
