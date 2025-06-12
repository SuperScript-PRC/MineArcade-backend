package packets

import (
	packets_arcade "MineArcade-backend/minearcade-server/protocol/packets/arcade"
	packet_define "MineArcade-backend/minearcade-server/protocol/packets/define"
	packets_general "MineArcade-backend/minearcade-server/protocol/packets/general"
	packets_lobby "MineArcade-backend/minearcade-server/protocol/packets/lobby"
)

// 客户端 -> 服务端 数据包
var ClientPool = map[uint32]func() ClientPacket{
	packet_define.IDClientHandshake:               func() ClientPacket { return &packets_general.ClientHandshake{} },
	packet_define.IDUDPConnection:                 func() ClientPacket { return &packets_general.UDPConnection{} },
	packet_define.IDClientLogin:                   func() ClientPacket { return &packets_general.ClientLogin{} },
	packet_define.IDDialLag:                       func() ClientPacket { return &packets_general.DialLag{} },
	packet_define.IDSimpleEvent:                   func() ClientPacket { return &packets_general.SimpleEvent{} },
	packet_define.IDStartGame:                     func() ClientPacket { return &packets_general.StartGame{} },
	packet_define.IDArcadeEntryRequest:            func() ClientPacket { return &packets_arcade.ArcadeEntryRequest{} },
	packet_define.IDArcadeMatchJoin:               func() ClientPacket { return &packets_arcade.ArcadeMatchJoin{} },
	packet_define.IDArcadeMatchEvent:              func() ClientPacket { return &packets_arcade.ArcadeMatchEvent{} },
	packet_define.IDArcadeExitGame:                func() ClientPacket { return &packets_arcade.ArcadeExitGame{} },
	packet_define.IDSimpleClientRequest:           func() ClientPacket { return &packets_general.SimpleClientRequest{} },
	packet_define.IDRankRequest:                   func() ClientPacket { return &packets_general.RankRequest{} },
	packet_define.IDPlaneFighterPlayerMove:        func() ClientPacket { return &packets_arcade.PlaneFighterPlayerMove{} },
	packet_define.IDPlaneFighterPlayerEvent:       func() ClientPacket { return &packets_arcade.PlaneFighterPlayerEvent{} },
	packet_define.IDPublicMineareaBlockEvent:      func() ClientPacket { return &packets_arcade.PublicMineareaBlockEvent{} },
	packet_define.IDPublicMineareaPlayerActorData: func() ClientPacket { return &packets_arcade.PublicMineareaPlayerActorData{} },
}

// 服务端 -> 客户端 数据包
var ServerPool = map[uint32]func() ServerPacket{
	packet_define.IDServerHandshake:               func() ServerPacket { return &packets_general.ServerHandshake{} },
	packet_define.IDClientLoginResp:               func() ServerPacket { return &packets_general.ClientLoginResp{} },
	packet_define.IDKickClient:                    func() ServerPacket { return &packets_general.KickClient{} },
	packet_define.IDDialLagResp:                   func() ServerPacket { return &packets_general.DialLagResp{} },
	packet_define.IDPlayerBasics:                  func() ServerPacket { return &packets_general.PlayerBasics{} },
	packet_define.IDBackpackResponse:              func() ServerPacket { return &packets_lobby.BackpackResponse{} },
	packet_define.IDSimpleEvent:                   func() ServerPacket { return &packets_general.SimpleEvent{} },
	packet_define.IDArcadeMatchJoinResp:           func() ServerPacket { return &packets_arcade.ArcadeMatchJoinResp{} },
	packet_define.IDArcadeMatchEvent:              func() ServerPacket { return &packets_arcade.ArcadeMatchEvent{} },
	packet_define.IDArcadeGameComplete:            func() ServerPacket { return &packets_arcade.ArcadeGameComplete{} },
	packet_define.IDPublicMineareaBlockEvent:      func() ServerPacket { return &packets_arcade.PublicMineareaBlockEvent{} },
	packet_define.IDPublicMineareaChunk:           func() ServerPacket { return &packets_arcade.PublicMineAreaChunk{} },
	packet_define.IDPublicMineareaPlayerActorData: func() ServerPacket { return &packets_arcade.PublicMineareaPlayerActorData{} },
	packet_define.IDArcadeEntryResponse:           func() ServerPacket { return &packets_arcade.ArcadeEntryResponse{} },
	packet_define.IDRankResponse:                  func() ServerPacket { return &packets_general.RankResponse{} },
	packet_define.IDPlaneFighterPlayerList:        func() ServerPacket { return &packets_arcade.PlaneFighterPlayerList{} },
	packet_define.IDPlaneFighterAddActor:          func() ServerPacket { return &packets_arcade.PlaneFighterAddActor{} },
	packet_define.IDPlaneFighterActorEvent:        func() ServerPacket { return &packets_arcade.PlaneFighterActorEvent{} },
	packet_define.IDPlaneFighterStage:             func() ServerPacket { return &packets_arcade.PlaneFighterStage{} },
	packet_define.IDPlaneFighterTimer:             func() ServerPacket { return &packets_arcade.PlaneFighterTimer{} },
	packet_define.IDPlaneFighterScores:            func() ServerPacket { return &packets_arcade.PlaneFighterScores{} },
	packet_define.IDPlaneFighterPlayerStatuses:    func() ServerPacket { return &packets_arcade.PlaneFighterPlayerStatuses{} },
}
