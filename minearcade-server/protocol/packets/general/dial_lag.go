package packets_general

import (
	"MineArcade-backend/minearcade-server/protocol"
	packet_define "MineArcade-backend/minearcade-server/protocol/packets/define"
)

// 由客户端发送, 请求检测服务端到客户端的延迟。
type DialLag struct {
	// 延迟检测数据包的 UUID
	dialUUID string
}

func (p *DialLag) ID() uint32 {
	return packet_define.IDDialLag
}

func (p *DialLag) NetType() int8 {
	return packet_define.UDPPacket
}

func (p *DialLag) Unmarshal(r *protocol.Reader) {
	r.StringUTF(&p.dialUUID)
}
