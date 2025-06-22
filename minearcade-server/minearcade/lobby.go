package minearcade

import (
	"MineArcade-backend/minearcade-server/clients"
	"MineArcade-backend/minearcade-server/clients/player_store"
	kickmsg "MineArcade-backend/minearcade-server/defines/kick_msg"
	packets_arcade "MineArcade-backend/minearcade-server/protocol/packets/arcade"
	packets_general "MineArcade-backend/minearcade-server/protocol/packets/general"
	"fmt"
	"log/slog"
)

func MainLobbyEntry(cli *clients.ArcadeClient) {
	store := player_store.ReadPlayerStore(cli.AuthInfo.UIDStr)
	cli.InitStoreInfo(store)
	cli.WritePacket(&packets_general.PlayerBasics{
		Nickname:   store.Nickname,
		UID:        cli.AuthInfo.UIDStr,
		Money:      store.Money,
		Power:      store.Power,
		Points:     store.Points,
		Level:      store.Level,
		Exp:        store.Exp,
		ExpUpgrade: store.ExpUpgrade,
	})
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
