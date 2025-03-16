package packets

import (
	"MineArcade-backend/protocol"
)

type PublicMineareaPlayerActorData struct {
	UUIDStr string
	X       float64
	Y       float64
	Action  uint8
}

func (ad *PublicMineareaPlayerActorData) ID() uint32 {
	return IDPublicMineareaPlayerActorData
}

func (ad *PublicMineareaPlayerActorData) Marshal(w *protocol.Writer) {
	w.StringUTF(ad.UUIDStr)
	w.Double(ad.X)
	w.Double(ad.Y)
	w.UInt8(ad.Action)
}

func (ad *PublicMineareaPlayerActorData) Unmarshal(r *protocol.Reader) {
	r.StringUTF(&ad.UUIDStr)
	r.Double(&ad.X)
	r.Double(&ad.Y)
	r.UInt8(&ad.Action)
}
