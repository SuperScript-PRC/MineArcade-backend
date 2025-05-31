package packets_general

import (
	"MineArcade-backend/minearcade-server/protocol"
	packet_define "MineArcade-backend/minearcade-server/protocol/packets/define"
)

// 由服务端返回, 检测服务端到客户端的延迟。
type DialLagResp struct {
	// 延迟检测数据包的 UUID
	dialUUID string
}

func (p *DialLagResp) ID() uint32 {
	return packet_define.IDDialLagResp
}

func (p *DialLagResp) NetType() int8 {
	return packet_define.UDPPacket
}

func (p *DialLagResp) Marshal(w *protocol.Writer) {
	w.StringUTF(p.dialUUID)
}
