package packets_general

import (
	"MineArcade-backend/minearcade-server/protocol"
	packet_define "MineArcade-backend/minearcade-server/protocol/packets/define"
)

// 简单事件, 如签到等
type SimpleEvent struct {
	// 事件类型
	EventType int32
	// 事件的附加数据
	EventData int32
}

func (p *SimpleEvent) ID() uint32 {
	return packet_define.IDSimpleEvent
}

func (p *SimpleEvent) NetType() int8 {
	return packet_define.TCPPacket
}

func (p *SimpleEvent) Marshal(w *protocol.Writer) {
	w.Int32(p.EventType)
	w.Int32(p.EventData)
}

func (p *SimpleEvent) Unmarshal(r *protocol.Reader) {
	r.Int32(&p.EventType)
	r.Int32(&p.EventData)
}
