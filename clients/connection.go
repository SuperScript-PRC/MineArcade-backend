package clients

import (
	"MineArcade-backend/defines"
	"MineArcade-backend/protocol/packets"
	"net"

	"github.com/pterm/pterm"
)

func HandleConnection(conn net.Conn) {
	defer conn.Close()
	cli := NetClient{Conn: conn}
	pks, err := cli.ReadPackets()
	if err != nil || len(pks) == 0 {
		if err != nil {
			pterm.Error.Println("读取数据包出错:", err)
		} else if len(pks) == 0 {
			pterm.Error.Println("没有读取到数据包")
		}
		return
	}
	like_handshake_pk := pks[0]
	if like_handshake_pk == nil {
		pterm.Error.Println("Read packet is nil")
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
		ServerMessage: "连接成功",
		ServerVersion: defines.MINEARCADE_VERSION,
	})
	pterm.Info.Println(cli.Conn.RemoteAddr().String() + ": 连接完成")
}
