package packets_arcade

import (
	"MineArcade-backend/minearcade-server/protocol"
	packet_define "MineArcade-backend/minearcade-server/protocol/packets/define"
	"MineArcade-backend/minearcade-server/protocol/ptypes/arcade_types"
)

type ArcadeMatchJoinResp struct {
	Success        bool
	Message        string
	CurrentPlayers []arcade_types.ArcadeMatchPlayer
}

func (p *ArcadeMatchJoinResp) ID() uint32 {
	return packet_define.IDArcadeMatchJoinResp
}

func (p *ArcadeMatchJoinResp) NetType() int8 {
	return packet_define.TCPPacket
}

func (p *ArcadeMatchJoinResp) Marshal(w *protocol.Writer) {
	w.Bool(p.Success)
	w.StringUTF(p.Message)
	protocol.WriteSlice(w, p.CurrentPlayers)
}
