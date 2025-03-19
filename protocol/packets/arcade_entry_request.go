package packets

import "MineArcade-backend/protocol"

type ArcadeEntryRequest struct {
	ArcadeGameType int8
	EntryID        string
	ResponseUUID   string
}

func (p *ArcadeEntryRequest) ID() uint32 {
	return IDArcadeEntryRequest
}

func (p *ArcadeEntryRequest) Unmarshal(r *protocol.Reader) {
	r.Int8(&p.ArcadeGameType)
	r.StringUTF(&p.EntryID)
	r.StringUTF(&p.ResponseUUID)
}
