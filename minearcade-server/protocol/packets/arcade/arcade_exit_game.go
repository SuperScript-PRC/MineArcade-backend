package packets_arcade

import (
	"MineArcade-backend/minearcade-server/protocol"
	packet_define "MineArcade-backend/minearcade-server/protocol/packets/define"
)

type ArcadeExitGame struct {
}

func (p *ArcadeExitGame) ID() uint32 {
	return packet_define.IDArcadeExitGame
}

func (p *ArcadeExitGame) NetType() int8 {
	return packet_define.TCPPacket
}

func (p *ArcadeExitGame) Unmarshal(w *protocol.Reader) {

}
