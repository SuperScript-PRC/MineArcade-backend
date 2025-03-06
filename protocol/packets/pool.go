package packets

const (
	IDClientLogin = iota
	IDClientLoginResp
	IDKickClient
	IDDialLag
	IDDialLagResp
	IDPlayerBasics
	IDBackpackResponse
	IDSimpleEvent
)

var ClientPool = map[uint32]func() ClientPacket{
	IDClientLogin: func() ClientPacket { return &ClientLogin{} },
	IDDialLag:     func() ClientPacket { return &DialLag{} },
	IDSimpleEvent: func() ClientPacket { return &SimpleEvent{} },
}

var ServerPool = map[uint32]func() ServerPacket{
	IDClientLoginResp:  func() ServerPacket { return &ClientLoginResp{} },
	IDDialLagResp:      func() ServerPacket { return &DialLagResp{} },
	IDPlayerBasics:     func() ServerPacket { return &PlayerBasics{} },
	IDKickClient:       func() ServerPacket { return &KickClient{} },
	IDBackpackResponse: func() ServerPacket { return &BackpackResponse{} },
	IDSimpleEvent:      func() ServerPacket { return &SimpleEvent{} },
}
