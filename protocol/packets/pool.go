package packets

const (
	IDClientHandshake = iota + 1
	IDServerHandshake
	IDClientLogin
	IDClientLoginResp
	IDKickClient
	IDDialLag
	IDDialLagResp
	IDPlayerBasics
	IDBackpackResponse
	IDSimpleEvent
)

// 客户端 -> 服务端 数据包
var ClientPool = map[uint32]func() ClientPacket{
	IDClientHandshake: func() ClientPacket { return &ClientHandshake{} },
	IDClientLogin:     func() ClientPacket { return &ClientLogin{} },
	IDDialLag:         func() ClientPacket { return &DialLag{} },
	IDSimpleEvent:     func() ClientPacket { return &SimpleEvent{} },
}

// 服务端 -> 客户端 数据包
var ServerPool = map[uint32]func() ServerPacket{
	IDClientLoginResp:  func() ServerPacket { return &ClientLoginResp{} },
	IDServerHandshake:  func() ServerPacket { return &ServerHandshake{} },
	IDKickClient:       func() ServerPacket { return &KickClient{} },
	IDDialLagResp:      func() ServerPacket { return &DialLagResp{} },
	IDPlayerBasics:     func() ServerPacket { return &PlayerBasics{} },
	IDBackpackResponse: func() ServerPacket { return &BackpackResponse{} },
	IDSimpleEvent:      func() ServerPacket { return &SimpleEvent{} },
}
