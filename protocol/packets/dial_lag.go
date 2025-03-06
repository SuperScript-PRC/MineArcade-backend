package packets

import "MineArcade-backend/protocol"

// 由客户端发送, 请求检测服务端到客户端的延迟。
type DialLag struct {
	// 延迟检测数据包的 UUID
	dialUUID string
}

func (p *DialLag) ID() uint32 {
	return IDDialLag
}

func (p *DialLag) Unmarshal(r *protocol.Reader) {
	r.StringUTF(&p.dialUUID)
}
