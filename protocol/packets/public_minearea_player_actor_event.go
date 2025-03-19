package packets

import (
	"MineArcade-backend/protocol"
)

const (
	MineAreaPlayerActionNone = iota
	MineAreaPlayerActionAddPlayer
	MineAreaPlayerActionRemovePlayer
)

type PublicMineareaPlayerActorData struct {
	Nickname string
	UIDStr   string
	X        float64
	Y        float64
	Action   int8
}

func (ad *PublicMineareaPlayerActorData) ID() uint32 {
	return IDPublicMineareaPlayerActorData
}

func (ad *PublicMineareaPlayerActorData) Marshal(w *protocol.Writer) {
	w.StringUTF(ad.Nickname)
	w.StringUTF(ad.UIDStr)
	w.Double(ad.X)
	w.Double(ad.Y)
	w.Int8(ad.Action)
}

func (ad *PublicMineareaPlayerActorData) Unmarshal(r *protocol.Reader) {
	r.StringUTF(&ad.Nickname)
	r.StringUTF(&ad.UIDStr)
	r.Double(&ad.X)
	r.Double(&ad.Y)
	r.Int8(&ad.Action)
}
