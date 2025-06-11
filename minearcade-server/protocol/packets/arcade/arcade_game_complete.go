package packets_arcade

import (
	"MineArcade-backend/minearcade-server/protocol"
	packet_define "MineArcade-backend/minearcade-server/protocol/packets/define"
)

type ArcadeGameComplete struct {
	TotalScore int32
}

func (p *ArcadeGameComplete) ID() uint32 {
	return packet_define.IDArcadeGameComplete
}

func (p *ArcadeGameComplete) NetType() int8 {
	return packet_define.TCPPacket
}

func (p *ArcadeGameComplete) Marshal(w *protocol.Writer) {
	w.Int32(p.TotalScore)
}
