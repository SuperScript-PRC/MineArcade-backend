package packets

import (
	"MineArcade-backend/protocol"
)

type PlayerActorData struct {
	UUIDStr string
	X       float64
	Y       float64
	Action  uint8
}

func (ad PlayerActorData) Marshal(w *protocol.Writer) {
	w.StringUTF(ad.UUIDStr)
	w.Double(ad.X)
	w.Double(ad.Y)
	w.UInt8(ad.Action)
}

func (ad *PlayerActorData) Unmarshal(r *protocol.Reader) {
	r.StringUTF(&ad.UUIDStr)
	r.Double(&ad.X)
	r.Double(&ad.Y)
	r.UInt8(&ad.Action)
}
