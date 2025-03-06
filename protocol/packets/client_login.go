package packets

import (
	"MineArcade-backend/protocol"
)

// 客户端登录请求
type ClientLogin struct {
	// 用户名
	Username string
	// 密码加盐后的 MD5 值
	Password string
}

func (p *ClientLogin) ID() uint32 {
	return IDClientLogin
}

func (p *ClientLogin) Unmarshal(r *protocol.Reader) {
	r.StringUTF(&p.Username)
	r.StringUTF(&p.Password)
}
