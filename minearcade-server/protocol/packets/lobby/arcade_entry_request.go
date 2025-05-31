package packets_lobby

import (
	"MineArcade-backend/minearcade-server/protocol"
	packet_define "MineArcade-backend/minearcade-server/protocol/packets/define"
)

type ArcadeEntryRequest struct {
	// 要加入的游戏的类型。
	ArcadeGameType int8
	// 游戏的 EntryID 数据，可以为房间号等。
	EntryID string
	// 请求的 UUID，会与 ResponseUUID 匹配。
	RequestUUID string
}

func (p *ArcadeEntryRequest) ID() uint32 {
	return packet_define.IDArcadeEntryRequest
}

func (p *ArcadeEntryRequest) NetType() int8 {
	return packet_define.TCPPacket
}

func (p *ArcadeEntryRequest) Unmarshal(r *protocol.Reader) {
	r.Int8(&p.ArcadeGameType)
	r.StringUTF(&p.EntryID)
	r.StringUTF(&p.RequestUUID)
}
