package packet_define

const (
	// General
	IDClientHandshake = iota + 1
	IDServerHandshake
	IDUDPConnection
	IDClientLogin
	IDClientLoginResp
	IDKickClient
	IDDialLag
	IDDialLagResp
	IDSimpleEvent
	IDSimpleClientRequest
	// Lobby
	IDPlayerBasics
	IDBackpackResponse
	IDRankRequest
	IDRankResponse
	IDWorldChat
	IDArcadeEntryRequest
	IDArcadeEntryResponse
	IDStartGame
	// Arcade:PublicMineArea
	IDPublicMineareaChunk
	IDPublicMineareaBlockEvent
	IDPublicMineareaPlayerActorData
	// Max
	MaxPacketID
)
