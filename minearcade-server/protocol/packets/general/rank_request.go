package packets_general

import (
	"MineArcade-backend/minearcade-server/protocol"
	packet_define "MineArcade-backend/minearcade-server/protocol/packets/define"
)

type RankRequest struct {
	// 请求的排行榜类型。每种街机游戏都有不同的排行榜类型。
	RankType int8
	// 请求的 UUID。会在 Response 中返回。
	RequestUUID string
}

func (p *RankRequest) ID() uint32 {
	return packet_define.IDRankRequest
}

func (p *RankRequest) NetType() int8 {
	return packet_define.TCPPacket
}

func (p *RankRequest) Unmarshal(r *protocol.Reader) {
	r.Int8(&p.RankType)
	r.StringUTF(&p.RequestUUID)
}
