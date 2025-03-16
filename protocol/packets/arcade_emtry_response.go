package packets

import "MineArcade-backend/protocol"

type ArcadeEntryResponse struct {
	ArcadeGameType int8
	ResponseUUID   string
	Success        bool
}

func (p *ArcadeEntryResponse) ID() uint32 {
	return IDArcadeEntryResponse
}

func (p *ArcadeEntryResponse) Marshal(w *protocol.Writer) {
	w.Int8(p.ArcadeGameType)
	w.StringUTF(p.ResponseUUID)
	w.Bool(p.Success)
}
