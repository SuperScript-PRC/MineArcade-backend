package packets_general

import (
	"MineArcade-backend/minearcade-server/protocol"
	packet_define "MineArcade-backend/minearcade-server/protocol/packets/define"
	"MineArcade-backend/minearcade-server/protocol/ptypes"
)

// 排行榜请求返回。
type RankResponse struct {
	// 排行榜数据。
	Ranks []ptypes.RankData
	// 请求的客户端在排行榜的信息。
	PlayerRank ptypes.RankData
}

func (p *RankResponse) ID() uint32 {
	return packet_define.IDRankResponse
}

func (p *RankResponse) NetType() int8 {
	return packet_define.TCPPacket
}

func (p *RankResponse) Marshal(w *protocol.Writer) {
	protocol.WriteSlice(w, p.Ranks)
	p.PlayerRank.Marshal(w)
}
