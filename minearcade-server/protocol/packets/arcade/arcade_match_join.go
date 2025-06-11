package packets_arcade

import (
	"MineArcade-backend/minearcade-server/protocol"
	packet_define "MineArcade-backend/minearcade-server/protocol/packets/define"
)

const (
	GameModeSolo = 0
	GameModeMutiple
)

type ArcadeMatchJoin struct {
	ArcadeGameType int8
	GameMode       int8
	RoomID         int32
}

func (p *ArcadeMatchJoin) ID() uint32 {
	return packet_define.IDArcadeMatchJoin
}

func (p *ArcadeMatchJoin) NetType() int8 {
	return packet_define.TCPPacket
}

func (p *ArcadeMatchJoin) Unmarshal(r *protocol.Reader) {
	r.Int8(&p.ArcadeGameType)
	r.Int8(&p.GameMode)
	r.Int32(&p.RoomID)
}
