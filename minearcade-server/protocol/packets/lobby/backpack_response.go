package packets_lobby

import (
	"MineArcade-backend/minearcade-server/protocol"
	packet_define "MineArcade-backend/minearcade-server/protocol/packets/define"
	"MineArcade-backend/minearcade-server/protocol/ptypes"
)

// 响应客户端的查询背包请求。
type BackpackResponse struct {
	// 背包物品列表
	Items []ptypes.Item
}

func (p *BackpackResponse) ID() uint32 {
	return packet_define.IDBackpackResponse
}

func (p *BackpackResponse) NetType() int8 {
	return packet_define.TCPPacket
}

func (p *BackpackResponse) Marshal(w *protocol.Writer) {
	protocol.WriteSlice(w, p.Items)
}
