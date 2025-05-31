package packets_general

import (
	"MineArcade-backend/minearcade-server/protocol"
	packet_define "MineArcade-backend/minearcade-server/protocol/packets/define"
)

// 简单客户端请求。
// 可以用于请求排行榜数据等。
type SimpleClientRequest struct {
	RequestType int32
	RequestUUID string
}

func (p *SimpleClientRequest) ID() uint32 {
	return packet_define.IDSimpleClientRequest
}

func (p *SimpleClientRequest) NetType() int8 {
	return packet_define.TCPPacket
}

func (p *SimpleClientRequest) Unmarshal(r *protocol.Reader) {
	r.Int32(&p.RequestType)
	r.StringUTF(&p.RequestUUID)
}
