package packets

import "MineArcade-backend/protocol"

// 客户端到服务端握手
type ClientHandshake struct {
	// 客户端版本
	ClientVersion int32
}

func (p *ClientHandshake) ID() uint32 {
	return IDClientHandshake
}

func (p *ClientHandshake) Unmarshal(r *protocol.Reader) {
	r.Int32(&p.ClientVersion)
}
