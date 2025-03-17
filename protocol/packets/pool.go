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
	IDPublicMineareaChunk
	IDPublicMineareaBlockEvent
	IDPublicMineareaPlayerActorData
	IDArcadeEntryRequest
	IDArcadeEntryResponse
)

// 客户端 -> 服务端 数据包
var ClientPool = map[uint32]func() ClientPacket{
	IDClientHandshake:               func() ClientPacket { return &ClientHandshake{} },
	IDClientLogin:                   func() ClientPacket { return &ClientLogin{} },
	IDDialLag:                       func() ClientPacket { return &DialLag{} },
	IDSimpleEvent:                   func() ClientPacket { return &SimpleEvent{} },
	IDPublicMineareaBlockEvent:      func() ClientPacket { return &PublicMineareaBlockEvent{} },
	IDPublicMineareaPlayerActorData: func() ClientPacket { return &PublicMineareaPlayerActorData{} },
	IDArcadeEntryRequest:            func() ClientPacket { return &ArcadeEntryRequest{} },
}

// 服务端 -> 客户端 数据包
var ServerPool = map[uint32]func() ServerPacket{
	IDServerHandshake:               func() ServerPacket { return &ServerHandshake{} },
	IDClientLoginResp:               func() ServerPacket { return &ClientLoginResp{} },
	IDKickClient:                    func() ServerPacket { return &KickClient{} },
	IDDialLagResp:                   func() ServerPacket { return &DialLagResp{} },
	IDPlayerBasics:                  func() ServerPacket { return &PlayerBasics{} },
	IDBackpackResponse:              func() ServerPacket { return &BackpackResponse{} },
	IDSimpleEvent:                   func() ServerPacket { return &SimpleEvent{} },
	IDPublicMineareaBlockEvent:      func() ServerPacket { return &PublicMineareaBlockEvent{} },
	IDPublicMineareaChunk:           func() ServerPacket { return &PublicMineAreaChunk{} },
	IDPublicMineareaPlayerActorData: func() ServerPacket { return &PublicMineareaPlayerActorData{} },
	IDArcadeEntryResponse:           func() ServerPacket { return &ArcadeEntryResponse{} },
}
