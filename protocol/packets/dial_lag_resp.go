package packets

import "MineArcade-backend/protocol"

// 由服务端返回, 检测服务端到客户端的延迟。
type DialLagResp struct {
	// 延迟检测数据包的 UUID
	dialUUID string
}

func (p *DialLagResp) ID() uint32 {
	return IDDialLagResp
}

func (p *DialLagResp) Marshal(w *protocol.Writer) {
	w.StringUTF(p.dialUUID)
}
