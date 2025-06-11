package packets_arcade

import (
	"MineArcade-backend/minearcade-server/protocol"
	packet_define "MineArcade-backend/minearcade-server/protocol/packets/define"
	"MineArcade-backend/minearcade-server/protocol/ptypes/arcade_types"
)

const (
	ArcadeMatchEventJoin = iota
	ArcadeMatchEventLeave
	ArcadeMatchEventAccept
	ArcadeMatchEventReady
)

type ArcadeMatchEvent struct {
	Action int8
	Player arcade_types.ArcadeMatchPlayer // Only available when S2C && Action != ArcadeMatchEventReady
}

func (p *ArcadeMatchEvent) ID() uint32 {
	return packet_define.IDArcadeMatchEvent
}

func (p *ArcadeMatchEvent) NetType() int8 {
	return packet_define.TCPPacket
}

func (p *ArcadeMatchEvent) Marshal(w *protocol.Writer) {
	w.Int8(p.Action)
	if p.Action != ArcadeMatchEventReady {
		p.Player.Marshal(w)
	}
}

func (p *ArcadeMatchEvent) Unmarshal(r *protocol.Reader) {
	r.Int8(&p.Action)
}
