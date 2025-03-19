package entry

import (
	"MineArcade-backend/arcade/public_minearea"
	"MineArcade-backend/clients"
	"MineArcade-backend/defines"
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
		pterm.Info.Printfln("等待玩家 IP=%v 加入游戏", cli.IPString)
		p, err := cli.ReadNextPacket()
		if err != nil {
			pterm.Error.Printfln("玩家 UID=%v 连接断开, 错误数据包: %v", cli.AuthInfo.UIDStr, err)
			cli.Kick("Broken packet")
			return
		}
		if pk, ok := p.(*packets.ArcadeEntryRequest); ok {
			switch pk.ArcadeGameType {
			case defines.GAMETYPE_PUBLIC_MINEAREA:
				pterm.Info.Printfln("玩家 UID=%v 准备加入游戏 Type=%v", cli.AuthInfo.UIDStr, pk.ArcadeGameType)
				ConfirmArcadeEntry(cli, pk, true)
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

func ConfirmArcadeEntry(cli *clients.NetClient, request_pk *packets.ArcadeEntryRequest, ok bool) {
	cli.WritePacket(&packets.ArcadeEntryResponse{ArcadeGameType: request_pk.ArcadeGameType, ResponseUUID: request_pk.ResponseUUID, Success: ok})
}
