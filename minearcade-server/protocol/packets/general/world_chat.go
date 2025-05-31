package packets_general

import (
	"MineArcade-backend/minearcade-server/protocol"
	packet_define "MineArcade-backend/minearcade-server/protocol/packets/define"
	"MineArcade-backend/minearcade-server/protocol/ptypes"
)

type WorldChat struct {
	ChatPlayer ptypes.GamePlayer
	Message    string
}

func (p *WorldChat) ID() uint32 {
	return packet_define.IDWorldChat
}

func (p *WorldChat) NetType() int8 {
	return packet_define.TCPPacket
}

func (p *WorldChat) Marshal(w *protocol.Writer) {
	p.ChatPlayer.Marshal(w)
	w.StringUTF(p.Message)
}
