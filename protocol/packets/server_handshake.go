package packets

import "MineArcade-backend/protocol"

// 客户端到服务端握手
type ServerHandshake struct {
	// 握手是否成功
	Success bool
	// 服务端版本
	ServerVersion int32
	// 服务端额外消息
	ServerMessage string
}

func (p *ServerHandshake) ID() uint32 {
	return IDServerHandshake
}

func (p *ServerHandshake) Marshal(w *protocol.Writer) {
	w.Bool(p.Success)
	w.Int32(p.ServerVersion)
	w.StringUTF(p.ServerMessage)
}
