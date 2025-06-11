package arcade_types

import "MineArcade-backend/minearcade-server/protocol"

type ArcadeMatchPlayer struct {
	Username string
	Nickname string
	UUID     string
}

func (p *ArcadeMatchPlayer) Marshal(w *protocol.Writer) {
	w.StringUTF(p.Username)
	w.StringUTF(p.Nickname)
	w.StringUTF(p.UUID)
}
