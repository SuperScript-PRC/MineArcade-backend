package minearcade

import (
	plane_fighter "MineArcade-backend/minearcade-server/arcade/plane_fighter"
	"MineArcade-backend/minearcade-server/arcade/public_minearea"
	"MineArcade-backend/minearcade-server/clients"
	"MineArcade-backend/minearcade-server/defines"
	"MineArcade-backend/minearcade-server/defines/kick_msg"
	packets_arcade "MineArcade-backend/minearcade-server/protocol/packets/arcade"
	"fmt"
	"log/slog"
)

func GoArcade(cli *clients.ArcadeClient, arcade_type int8, request_uuid string) {
	slog.Info(fmt.Sprintf("玩家 %s 准备加入游戏 Type=%v", cli.FormatNameWithUUID(), arcade_type))
	switch arcade_type {
	case defines.GAMETYPE_PUBLIC_MINEAREA:
		ConfirmArcadeEntry(cli, arcade_type, request_uuid, true)
		public_minearea.PlayerEntry(cli)
	case defines.GAMETYPE_PLANE_FIGHTER:
		ConfirmArcadeEntry(cli, arcade_type, request_uuid, true)
		if room, ok := plane_fighter.PlayerPrejoinEntry(cli); ok {
			plane_fighter.PlayerEntry(cli, room)
		}
	default:
		cli.Kick(kick_msg.INVALID_PACKET)
	}
}

func ConfirmArcadeEntry(cli *clients.ArcadeClient, arcade_type int8, request_uuid string, ok bool) {
	cli.WritePacket(&packets_arcade.ArcadeEntryResponse{ArcadeGameType: arcade_type, ResponseUUID: request_uuid, Success: ok})
}
