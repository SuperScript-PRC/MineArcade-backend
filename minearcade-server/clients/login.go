package clients

import (
	"MineArcade-backend/minearcade-server/clients/accounts"
	"MineArcade-backend/minearcade-server/clients/player_store"
	"MineArcade-backend/minearcade-server/configs"
	"MineArcade-backend/minearcade-server/defines"
	"MineArcade-backend/minearcade-server/defines/kick_msg"
	packet_define "MineArcade-backend/minearcade-server/protocol/packets/define"
	packets_general "MineArcade-backend/minearcade-server/protocol/packets/general"

	"github.com/google/uuid"
	"github.com/pterm/pterm"
)

func ClientLogin(cli *ArcadeClient) bool {
	like_handshake_pk, err := cli.NextPacket()
	if err != nil {
		pterm.Error.Println("读取数据包出错:", err)
		return false
	}
	handshake_pk, ok := like_handshake_pk.(*packets_general.ClientHandshake)
	if !ok {
		cli.WritePacket(&packets_general.ServerHandshake{
			Success:       false,
			ServerMessage: "Handshake packet ERROR",
			ServerVersion: defines.MINEARCADE_VERSION,
		})
		pterm.Error.Printfln("握手失败: 客户端握手包错误: ID=%v", like_handshake_pk.ID())
		return false
	}
	verify_token := uuid.New().String()
	if handshake_pk.ClientVersion < defines.MINEARCADE_VERSION {
		cli.WritePacket(&packets_general.ServerHandshake{
			Success:       false,
			ServerMessage: "客户端版本过低",
			ServerVersion: defines.MINEARCADE_VERSION,
			VerifyToken:   "",
		})
		pterm.Error.Printfln("握手失败: 客户端版本 %v < %v", handshake_pk.ClientVersion, defines.MINEARCADE_VERSION)
		return false
	} else if handshake_pk.ClientVersion > defines.MINEARCADE_VERSION {
		cli.WritePacket(&packets_general.ServerHandshake{
			Success:       false,
			ServerMessage: "服务端版本过低",
			ServerVersion: defines.MINEARCADE_VERSION,
			VerifyToken:   "",
		})
		pterm.Error.Printfln("握手失败: 客户端版本 %v > %v", handshake_pk.ClientVersion, defines.MINEARCADE_VERSION)
		return false
	}
	// TCP Handshake OK
	cli.WritePacket(&packets_general.ServerHandshake{
		Success:       true,
		ServerMessage: "",
		ServerVersion: defines.MINEARCADE_VERSION,
		VerifyToken:   verify_token,
	})
	_recv_pk, ok := cli.WaitForPacket(packet_define.IDUDPConnection, configs.UDP_CONNECTION_TIMEOUT)
	if !ok {
		pterm.Warning.Printfln("客户端 %s UDPConnection 超时", cli.TCPAddr)
		cli.Kick(kick_msg.UDP_CONNECTION_TIMEOUT)
		return false
	}
	recv_pk, ok := _recv_pk.(*packets_general.UDPConnection)
	if !ok {
		cli.Kick(kick_msg.INVALID_PACKET)
		return false
	} else if recv_pk.VerifyToken != verify_token {
		cli.Kick("Invalid token")
		return false
	}
	pterm.Success.Printfln("%v 握手成功", cli.TCPAddr.String())
	// wait login
	for {
		// password trial
		like_login_pk, err := cli.NextPacket()
		if err != nil {
			pterm.Error.Println("读取数据包出错:", err)
			return false
		}
		login_pk, ok := like_login_pk.(*packets_general.ClientLogin)
		if !ok {
			cli.WritePacket(&packets_general.ClientLoginResp{
				Success:    false,
				Message:    "Login packet ERROR",
				StatusCode: 1,
			})
			//continue
			pterm.Error.Printfln("%v 登录失败: 客户端登录包错误: ID=%v", cli.TCPAddr.String(), like_login_pk.ID())
			// if pk1, ok := like_login_pk.(*packets_general.ClientHandshake); ok {
			// 	println("ID1:", pk1.ClientVersion)
			// }
			// if pk2, err := cli.NextPacket(); err == nil {
			// 	println("ID2:", pk2.ID())
			// }
			return false
		}
		account_ok, reason := accounts.IsAccountOK(login_pk.Username, login_pk.Password)
		if !account_ok {
			cli.WritePacket(&packets_general.ClientLoginResp{
				Success:    false,
				Message:    reason,
				StatusCode: 1,
			})
			pterm.Warning.Printfln("%v 登录失败: 账号或密码错误: %v, %v", cli.TCPAddr.String(), login_pk.Username, login_pk.Password)
		} else {
			cli.WritePacket(&packets_general.ClientLoginResp{
				Success:    true,
				Message:    "登录成功",
				StatusCode: 0,
			})
			userinfo, ok := accounts.GetUserAuthInfo(login_pk.Username)
			if !ok {
				panic("Auth failed?? Shouldn't be happened")
			}
			cli.InitAuthInfo(userinfo)
			pterm.Success.Printfln("%v 登录成功, 账号: %v, UID: %v, 昵称: %v", cli.TCPAddr.String(), cli.AuthInfo.AccountName, cli.AuthInfo.UIDStr, cli.AuthInfo.Nickname)
			break
		}
	}
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
	return true
}
