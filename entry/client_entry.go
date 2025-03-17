package entry

import (
	"MineArcade-backend/arcade/public_minearea"
	"MineArcade-backend/clients"
	"MineArcade-backend/protocol/packets"
	"net"

	"github.com/pterm/pterm"
)

func HandleClientConnection(conn net.Conn) {
	cli, ok := clients.HandleConnection(conn)
	if !ok {
		return
	}
	ClientEntry(cli)
}

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
				pterm.Info.Printfln("玩家 UUID=%v 准备加入游戏 Type=%v", cli.AuthInfo.UUIDStr, pk.ArcadeGameType)
				cli.WritePacket(&packets.ArcadeEntryResponse{
					ArcadeGameType: pk.ArcadeGameType,
					ResponseUUID:   pk.RequestUUID,
					Success:        true,
				})
				go public_minearea.PlayerEntry(cli)
				return
			default:
				cli.Kick("Lobby: 不合法的数据包")
			}
		} else {
			pterm.Warning.Printfln("客户端发送不正确数据包: %d", p.ID())
			cli.Kick("Lobby: 不合法的数据包")
		}
	}
}

func ConfirmArcadeEntry(cli *clients.NetClient, request_pk *packets.ArcadeEntryRequest, ok bool) *packets.ArcadeEntryResponse {
	return &packets.ArcadeEntryResponse{ArcadeGameType: request_pk.ArcadeGameType, ResponseUUID: request_pk.RequestUUID, Success: ok}
}
