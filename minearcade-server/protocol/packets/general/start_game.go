package packets_general

import (
	"MineArcade-backend/minearcade-server/protocol"
	packet_define "MineArcade-backend/minearcade-server/protocol/packets/define"
)

type StartGame struct {
	ArcadeGameType int8
	EntryID        string
}

func (p *StartGame) ID() uint32 {
	return packet_define.IDStartGame
}

func (p *StartGame) NetType() int8 {
	return packet_define.TCPPacket
}

func (p *StartGame) Unmarshal(r *protocol.Reader) {
	r.Int8(&p.ArcadeGameType)
	r.StringUTF(&p.EntryID)
}
