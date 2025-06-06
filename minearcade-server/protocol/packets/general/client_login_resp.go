package packets_general

import (
	"MineArcade-backend/minearcade-server/protocol"
	packet_define "MineArcade-backend/minearcade-server/protocol/packets/define"
)

const (
	LoginRespSuccess = iota
	LoginRespAccountNotFound
	LoginRespWrongPassword
	LoginRespIsBanning
	LoginRespServerIsFixing
)

// 客户端登录请求返回
type ClientLoginResp struct {
	// 是否成功
	Success bool
	// 如果不成功, 登录失败的消息
	Message string
	// 登录状态码
	StatusCode int8
}

func (p *ClientLoginResp) ID() uint32 {
	return packet_define.IDClientLoginResp
}

func (p *ClientLoginResp) NetType() int8 {
	return packet_define.TCPPacket
}

func (p *ClientLoginResp) Marshal(w *protocol.Writer) {
	w.Bool(p.Success)
	w.StringUTF(p.Message)
	w.Int8(p.StatusCode)
}
