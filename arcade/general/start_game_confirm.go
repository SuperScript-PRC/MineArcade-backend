package general

import (
	"MineArcade-backend/clients"
	"MineArcade-backend/defines"
	"MineArcade-backend/protocol/packets"
)

func ConfirmStartGame(cli *clients.NetClient, arcadeGameType int8) bool {
	pk, err := cli.ReadNextPacket()
	if err != nil {
		cli.Kick("Broken packet")
		return false
	}
	start_game_pk, ok := pk.(*packets.StartGame)
	if !ok {
		cli.Kick("Need StartGame")
		return false
	}
	if start_game_pk.ArcadeGameType != defines.GAMETYPE_PUBLIC_MINEAREA {
		cli.Kick("Invalid StartGame")
		return false
	}
	return true
}
