package packets

import "MineArcade-backend/protocol"

//  将 客户端->服务端数据包 与 服务端->客户端数据包 分开管理

type ClientPacket interface {
	ID() uint32
	Unmarshal(r *protocol.Reader)
}

type ServerPacket interface {
	ID() uint32
	Marshal(w *protocol.Writer)
}
