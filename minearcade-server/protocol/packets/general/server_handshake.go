package packets_general

import (
	"MineArcade-backend/minearcade-server/protocol"
	packet_define "MineArcade-backend/minearcade-server/protocol/packets/define"
)

// 客户端到服务端握手
type ServerHandshake struct {
	// 握手是否成功
	Success bool
	// 服务端版本
	ServerVersion int32
	// 服务端额外消息
	ServerMessage string
	// 用于 UDPConnection 验证的 Token
	VerifyToken string
}

func (p *ServerHandshake) ID() uint32 {
	return packet_define.IDServerHandshake
}

func (p *ServerHandshake) NetType() int8 {
	return packet_define.TCPPacket
}

func (p *ServerHandshake) Marshal(w *protocol.Writer) {
	w.Bool(p.Success)
	w.Int32(p.ServerVersion)
	w.StringUTF(p.ServerMessage)
	w.StringUTF(p.VerifyToken)
}
