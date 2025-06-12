package packets_arcade

import (
	"MineArcade-backend/minearcade-server/protocol"
	packet_define "MineArcade-backend/minearcade-server/protocol/packets/define"
	"MineArcade-backend/minearcade-server/protocol/ptypes/arcade_types"
)

type ArcadeGameComplete struct {
	Win          bool
	TotalScore   int32
	ScoreDetails []arcade_types.ArcadeGameScoreDetail
}

func (p *ArcadeGameComplete) ID() uint32 {
	return packet_define.IDArcadeGameComplete
}

func (p *ArcadeGameComplete) NetType() int8 {
	return packet_define.TCPPacket
}

func (p *ArcadeGameComplete) Marshal(w *protocol.Writer) {
	w.Bool(p.Win)
	w.Int32(p.TotalScore)
	protocol.WriteSlice(w, p.ScoreDetails)
}
