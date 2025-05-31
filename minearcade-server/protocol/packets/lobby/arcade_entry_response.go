package packets_lobby

import (
	"MineArcade-backend/minearcade-server/protocol"
	packet_define "MineArcade-backend/minearcade-server/protocol/packets/define"
)

type ArcadeEntryResponse struct {
	// 此游戏的游戏类型
	ArcadeGameType int8
	// 回应的 UUID，对应着客户端发送的 ArcadeEntryRequest 中的 RequestUUID
	ResponseUUID string
	// 请求加入游戏是否成功
	Success bool
}

func (p *ArcadeEntryResponse) ID() uint32 {
	return packet_define.IDArcadeEntryResponse
}

func (p *ArcadeEntryResponse) NetType() int8 {
	return packet_define.TCPPacket
}

func (p *ArcadeEntryResponse) Marshal(w *protocol.Writer) {
	w.Int8(p.ArcadeGameType)
	w.StringUTF(p.ResponseUUID)
	w.Bool(p.Success)
}
