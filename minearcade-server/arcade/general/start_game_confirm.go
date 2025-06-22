package packets_general

import (
	clients "MineArcade-backend/minearcade-server/clients"
	"MineArcade-backend/minearcade-server/defines"
	"MineArcade-backend/minearcade-server/defines/kick_msg"
	packets_general "MineArcade-backend/minearcade-server/protocol/packets/general"
)

func ConfirmStartGame(cli *clients.ArcadeClient, arcadeGameType int8) bool {
	pk, err := cli.NextPacket()
	if err != nil {
		cli.Kick(kick_msg.BROKEN_PACKET)
		return false
	}
	start_game_pk, ok := pk.(*packets_general.StartGame)
	if !ok {
		// println("invalid:", pk.ID())
		cli.Kick("Need StartGame")
		return false
	}
	if start_game_pk.ArcadeGameType != defines.GAMETYPE_PUBLIC_MINEAREA {
		cli.Kick("Invalid StartGame")
		return false
	}
	return true
}
