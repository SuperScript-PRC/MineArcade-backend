package minearcade

import (
	"MineArcade-backend/minearcade-server/clients"
	kickmsg "MineArcade-backend/minearcade-server/defines/kick_msg"
	packets_lobby "MineArcade-backend/minearcade-server/protocol/packets/lobby"

	"github.com/pterm/pterm"
)

func MainLobbyEntry(cli *clients.ArcadeClient) {
	for {
		pterm.Info.Printfln("等待玩家 %s 加入游戏", cli.FormatNameWithUUID())
		p, err := cli.NextPacket()
		if err != nil {
			pterm.Error.Printfln("玩家 %s 连接断开, 错误: %v", cli.FormatNameWithUUID(), err)
			cli.Kick(kickmsg.BROKEN_PACKET)
			return
		}
		if pk, ok := p.(*packets_lobby.ArcadeEntryRequest); ok {
			GoArcade(cli, pk.ArcadeGameType, pk.RequestUUID)
		} else {
			pterm.Warning.Printfln("客户端发送不正确数据包: %d", p.ID())
			cli.Kick(kickmsg.INVALID_PACKET)
		}
	}
}
