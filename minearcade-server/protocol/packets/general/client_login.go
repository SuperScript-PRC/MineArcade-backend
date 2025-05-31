package packets_general

import (
	"MineArcade-backend/minearcade-server/protocol"
	packet_define "MineArcade-backend/minearcade-server/protocol/packets/define"
)

// 客户端登录请求
type ClientLogin struct {
	// 用户名
	Username string
	// 密码加盐后的 MD5 值
	Password string
}

func (p *ClientLogin) ID() uint32 {
	return packet_define.IDClientLogin
}

func (p *ClientLogin) NetType() int8 {
	return packet_define.TCPPacket
}

func (p *ClientLogin) Unmarshal(r *protocol.Reader) {
	r.StringUTF(&p.Username)
	r.StringUTF(&p.Password)
}
