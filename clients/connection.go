package clients

import (
	"MineArcade-backend/clients/accounts"
	"MineArcade-backend/clients/player_store"
	"MineArcade-backend/defines"
	"MineArcade-backend/protocol/packets"
	"net"

	"github.com/pterm/pterm"
)

func HandleConnection(conn net.Conn) {
	defer conn.Close()
	cli := NetClient{Conn: conn, IPString: conn.RemoteAddr().String()}
	like_handshake_pk, err := cli.ReadNextPacket()
	if err != nil {
		pterm.Error.Println("读取数据包出错:", err)
		return
	}
	handshake_pk, ok := like_handshake_pk.(*packets.ClientHandshake)
	if !ok {
		cli.WritePacket(&packets.ServerHandshake{
			Success:       false,
			ServerMessage: "Handshake packet ERROR",
			ServerVersion: defines.MINEARCADE_VERSION,
		})
		pterm.Error.Printfln("握手失败: 客户端握手包错误: ID=%v", like_handshake_pk.ID())
		return
	}
	if handshake_pk.ClientVersion < defines.MINEARCADE_VERSION {
		cli.WritePacket(&packets.ServerHandshake{
			Success:       false,
			ServerMessage: "客户端版本过低",
			ServerVersion: defines.MINEARCADE_VERSION,
		})
		pterm.Error.Printfln("握手失败: 客户端版本 %v < %v", handshake_pk.ClientVersion, defines.MINEARCADE_VERSION)
		return
	} else if handshake_pk.ClientVersion > defines.MINEARCADE_VERSION {
		cli.WritePacket(&packets.ServerHandshake{
			Success:       false,
			ServerMessage: "服务端版本过低",
			ServerVersion: defines.MINEARCADE_VERSION,
		})
		pterm.Error.Printfln("握手失败: 客户端版本 %v > %v", handshake_pk.ClientVersion, defines.MINEARCADE_VERSION)
		return
	}
	cli.WritePacket(&packets.ServerHandshake{
		Success:       true,
		ServerMessage: "",
		ServerVersion: defines.MINEARCADE_VERSION,
	})
	pterm.Success.Printfln("%v 握手成功", cli.IPString)
	// wait login
	for {
		// password trial
		like_login_pk, err := cli.ReadNextPacket()
		if err != nil {
			pterm.Error.Println("读取数据包出错:", err)
			return
		}
		login_pk, ok := like_login_pk.(*packets.ClientLogin)
		if !ok {
			cli.WritePacket(&packets.ServerHandshake{
				Success:       false,
				ServerMessage: "Login packet ERROR",
				ServerVersion: defines.MINEARCADE_VERSION,
			})
			pterm.Error.Printfln("%v 登录失败: 客户端登录包错误: ID=%v", cli.IPString, like_login_pk.ID())
			return
		}
		accountant_ok, reason := accounts.IsAccountOK(login_pk.Username, login_pk.Password)
		if !accountant_ok {
			cli.WritePacket(&packets.ClientLoginResp{
				Success:    false,
				Message:    reason,
				StatusCode: 1,
			})
			pterm.Warning.Printfln("%v 登录失败: 账号或密码错误: %v, %v", cli.IPString, login_pk.Username, login_pk.Password)
		} else {
			pterm.Success.Printfln("%v 登录成功", cli.IPString)
			cli.WritePacket(&packets.ClientLoginResp{
				Success:    true,
				Message:    "登录成功",
				StatusCode: 0,
			})
			userinfo, ok := accounts.GetUserAuthInfo(login_pk.Username)
			if !ok {
				panic("Auth failed?? Shouldn't be happened")
			}
			cli.InitAuthInfo(userinfo)
			break
		}
	}
	store := player_store.ReadPlayerStore(cli.AuthInfo.UUIDStr)
	cli.InitStoreInfo(store)
	cli.WritePacket(&packets.PlayerBasics{
		Nickname:   store.Nickname,
		UUID:       cli.AuthInfo.UUIDStr,
		Money:      store.Money,
		Power:      store.Power,
		Points:     store.Points,
		Level:      store.Level,
		Exp:        store.Exp,
		ExpUpgrade: store.ExpUpgrade,
	})
	cli.ReadNextPacket()
	pterm.Info.Println(cli.IPString + ": 连接完成")
}
