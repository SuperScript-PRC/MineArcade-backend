package packets_general

import (
	"MineArcade-backend/minearcade-server/protocol"
	packet_define "MineArcade-backend/minearcade-server/protocol/packets/define"
)

// 客户端到服务端握手
type ClientHandshake struct {
	// 客户端版本
	ClientVersion int32
}

func (p *ClientHandshake) ID() uint32 {
	return packet_define.IDClientHandshake
}

func (p *ClientHandshake) NetType() int8 {
	return packet_define.TCPPacket
}

func (p *ClientHandshake) Unmarshal(r *protocol.Reader) {
	r.Int32(&p.ClientVersion)
}
