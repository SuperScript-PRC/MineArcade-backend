package entry

import (
	"MineArcade-backend/arcade/public_minearea"
	"MineArcade-backend/clients"
	"MineArcade-backend/protocol/packets"
)

func ClientEntry(cli *clients.NetClient) {
	MainLobbyEntry(cli)
}

func MainLobbyEntry(cli *clients.NetClient) {
	for {
		p, err := cli.ReadNextPacket()
		if err != nil {
			cli.Kick("Broken packet")
			return
		}
		if pk, ok := p.(*packets.ArcadeEntryRequest); ok {
			switch pk.ArcadeGameType {
			case GAMETYPE_PUBLIC_MINEAREA:
				go public_minearea.PlayerEntry(cli)
			}
		}
	}
}

func ConfirmArcadeEntry(cli *clients.NetClient, request_pk *packets.ArcadeEntryRequest, ok bool) *packets.ArcadeEntryResponse {
	return &packets.ArcadeEntryResponse{ArcadeGameType: request_pk.ArcadeGameType, ResponseUUID: request_pk.RequestUUID, Success: ok}
}
