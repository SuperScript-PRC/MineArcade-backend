package packets_general

import (
	"MineArcade-backend/minearcade-server/protocol"
	packet_define "MineArcade-backend/minearcade-server/protocol/packets/define"
)

const (
	KickClientBanned = iota
	ServerDown
	ServerFixing
)

// 从服务器踢出一个客户端。
type KickClient struct {
	// 踢出客户端显示的信息
	Message string
	// 踢出的状态码
	StatusCode int8
}

func (p *KickClient) ID() uint32 {
	return packet_define.IDKickClient
}

func (p *KickClient) NetType() int8 {
	return packet_define.TCPPacket
}

func (p *KickClient) Marshal(w *protocol.Writer) {
	w.StringUTF(p.Message)
	w.Int8(p.StatusCode)
}
