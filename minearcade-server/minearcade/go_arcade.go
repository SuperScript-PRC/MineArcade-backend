package minearcade

import (
	"MineArcade-backend/minearcade-server/arcade/public_minearea"
	"MineArcade-backend/minearcade-server/clients"
	"MineArcade-backend/minearcade-server/defines"
	"MineArcade-backend/minearcade-server/defines/kick_msg"
	packets_lobby "MineArcade-backend/minearcade-server/protocol/packets/lobby"

	"github.com/pterm/pterm"
)

func GoArcade(cli *clients.ArcadeClient, arcade_type int8, request_uuid string) {
	pterm.Info.Printfln("玩家 %s 准备加入游戏 Type=%v", cli.FormatNameWithUUID(), arcade_type)
	switch arcade_type {
	case defines.GAMETYPE_PUBLIC_MINEAREA:
		ConfirmArcadeEntry(cli, arcade_type, request_uuid, true)
		public_minearea.PlayerEntry(cli)
	default:
		cli.Kick(kick_msg.INVALID_PACKET)
	}
}

func ConfirmArcadeEntry(cli *clients.ArcadeClient, arcade_type int8, request_uuid string, ok bool) {
	cli.WritePacket(&packets_lobby.ArcadeEntryResponse{ArcadeGameType: arcade_type, ResponseUUID: request_uuid, Success: ok})
}
