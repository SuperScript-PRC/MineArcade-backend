package packets

import "MineArcade-backend/protocol"

type StartGame struct {
	ArcadeGameType int8
	EntryID        string
}

func (p *StartGame) ID() uint32 {
	return IDStartGame
}

func (p *StartGame) Unmarshal(r *protocol.Reader) {
	r.Int8(&p.ArcadeGameType)
	r.StringUTF(&p.EntryID)
}
