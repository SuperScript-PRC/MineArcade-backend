package packets

import "MineArcade-backend/protocol"

// 简单事件, 如签到等
type SimpleEvent struct {
	// 事件类型
	EventType int32
	// 事件的附加数据
	EventData int32
}

func (p *SimpleEvent) ID() uint32 {
	return IDSimpleEvent
}

func (p *SimpleEvent) Marshal(w *protocol.Writer) {
	w.Int32(p.EventType)
	w.Int32(p.EventData)
}

func (p *SimpleEvent) Unmarshal(r *protocol.Reader) {
	r.Int32(&p.EventType)
	r.Int32(&p.EventData)
}
