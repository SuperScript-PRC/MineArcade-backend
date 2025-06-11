package plane_fighter

import (
	"MineArcade-backend/minearcade-server/arcade/plane_fighter/define"
	"MineArcade-backend/minearcade-server/clients"
	"MineArcade-backend/minearcade-server/defines"
	"MineArcade-backend/minearcade-server/defines/kick_msg"
	packets_arcade "MineArcade-backend/minearcade-server/protocol/packets/arcade"
	packet_define "MineArcade-backend/minearcade-server/protocol/packets/define"
	"MineArcade-backend/minearcade-server/protocol/ptypes/arcade_types"
	"time"
)

func PlayerPrejoinEntry(cli *clients.ArcadeClient) (room *PlaneFighterRoom, join_ok bool) {
	_pk, err := cli.NextPacket()
	if err != nil {
		cli.Kick(kick_msg.BROKEN_PACKET)
		return
	}
	pk, ok := _pk.(*packets_arcade.ArcadeMatchJoin)
	if !ok {
		cli.Kick(kick_msg.INVALID_PACKET)
		return
	} else if pk.ArcadeGameType != defines.GAMETYPE_PLANE_FIGHTER {
		cli.WritePacket(&packets_arcade.ArcadeMatchJoinResp{
			Success:        false,
			Message:        "无效的游戏类型",
			CurrentPlayers: []arcade_types.ArcadeMatchPlayer{},
		})
		return
	}
	cli.WritePacket(&packets_arcade.ArcadeMatchJoinResp{
		Success:        true,
		Message:        "",
		CurrentPlayers: []arcade_types.ArcadeMatchPlayer{},
	})
	if pk.GameMode == packets_arcade.GameModeSolo {
		room = NewRoomWithRandomRoomID(1)
	} else {
		_, room = GetAvailRoom()
	}
	go room.ActiveMatcher()
	room.AddClient(cli)
	join_ok = true
	return
}
func PlayerEntry(cli *clients.ArcadeClient, room *PlaneFighterRoom) {
	stg := room.Stage
	room.WaitMatchReady(cli)
	room.Stage.AddPlayer(NewPlayer(room.GetClientRuntimeID(cli)))
	_, getted := cli.WaitForPacket(packet_define.IDStartGame, time.Second*10)
	if !getted {
		// todo: StartGame timeout
		// do sth?
		println("timeout (todoooo)")
	}
	room.AddGameReadyCount()
	room.WaitGameReady()
	cli.WritePacket(room.MakePlayerList())
	for {
		_pk, err, interrupted := cli.NextPacketWithInterrupt(room.Closed)
		if err != nil {
			cli.Kick(kick_msg.BROKEN_PACKET)
			return
		} else if interrupted {
			return
		}
		switch pk := _pk.(type) {
		case *packets_arcade.PlaneFighterPlayerMove:
			stg.PlayerSyncPosition(room.GetClientRuntimeID(cli), pk.X, pk.Y)
		case *packets_arcade.PlaneFighterPlayerEvent:
			switch pk.EventID {
			case packets_arcade.PFPEventStartFire:
				stg.PlayerStartFire(room.GetClientRuntimeID(cli))
			case packets_arcade.PFPEventStopFire:
				stg.PlayerStopFire(room.GetClientRuntimeID(cli))
			}
		}
	}
}

func RoomEntry(room *PlaneFighterRoom) {
	stage := room.Stage
	for {
		print(".")
		stage.RunTick()
		room.SendStageEvents()
		room.SendStageNewActors()
		room.SendAddScore()
		room.SendStage()
		time.Sleep(define.SLEEP_TIME)
		if stage.isEnded() {
			break
		}
		if stage.Ticks%define.TPS == 0 {
			room.broadcastPacket(&packets_arcade.PlaneFighterTimer{SecondsLeft: (stage.TicksLeft / define.TPS)})
		}
		if !room.CheckAllOnline() {
			break
		}
	}
}
