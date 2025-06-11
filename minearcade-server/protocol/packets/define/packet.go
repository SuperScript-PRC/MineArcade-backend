package packet_define

import "MineArcade-backend/minearcade-server/protocol"

//  将 客户端->服务端数据包 与 服务端->客户端数据包 分开管理

// 客户端 -> 服务端
type ClientPacket interface {
	ID() uint32
	Unmarshal(r *protocol.Reader)
	NetType() int8
}

// 服务端 -> 客户端
type ServerPacket interface {
	ID() uint32
	NetType() int8
	Marshal(w *protocol.Writer)
}
