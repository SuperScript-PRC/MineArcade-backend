package plane_fighter

import (
	"MineArcade-backend/minearcade-server/arcade/plane_fighter/define"
	"MineArcade-backend/minearcade-server/clients"
	"MineArcade-backend/minearcade-server/protocol/packets"
	packets_arcade "MineArcade-backend/minearcade-server/protocol/packets/arcade"
	"MineArcade-backend/minearcade-server/protocol/ptypes/arcade_types"
	"fmt"

	"github.com/google/uuid"
)

type RoomEvent struct {
	JoinEvent         chan *clients.ArcadeClient
	MatchReady        chan bool
	MatchAlreadyReady bool
	GameReady         chan bool
	GameAlreadyReady  bool
	GameReadyCount    int
}

type PlaneFighterRoom struct {
	RoomEvents          *RoomEvent
	RoomID              string
	MaxPlayer           int
	Clients             []*clients.ArcadeClient
	ClientUID2RuntimeID map[string]int32
	Stage               *PlaneFighterStage
	Closed              chan bool
}

func NewRoom(roomID string, maxPlayer int) *PlaneFighterRoom {
	return &PlaneFighterRoom{
		RoomEvents:          NewRoomEvent(),
		RoomID:              roomID,
		MaxPlayer:           maxPlayer,
		ClientUID2RuntimeID: make(map[string]int32),
		Stage:               NewStage(),
	}
}

func NewRoomEvent() *RoomEvent {
	return &RoomEvent{
		MatchReady:        make(chan bool, 8),
		JoinEvent:         make(chan *clients.ArcadeClient, 64),
		GameReady:         make(chan bool, 8),
		GameAlreadyReady:  false,
		MatchAlreadyReady: false,
		GameReadyCount:    0,
	}
}

func NewRoomWithRandomRoomID(maxPlayer int) *PlaneFighterRoom {
	ud, err := uuid.NewUUID()
	if err != nil {
		panic(fmt.Errorf("RoomUD error: %v", err))
	}
	return &PlaneFighterRoom{
		RoomEvents:          NewRoomEvent(),
		RoomID:              ud.String(),
		MaxPlayer:           maxPlayer,
		ClientUID2RuntimeID: make(map[string]int32),
		Stage:               NewStage(),
	}
}

func (r *PlaneFighterRoom) ActiveMatcher() {
	if r.RoomEvents.MatchAlreadyReady {
		return
	}
	for range r.RoomEvents.JoinEvent {
		if len(r.Clients) >= r.MaxPlayer {
			if r.MaxPlayer == 1 {
				r.RoomEvents.MatchAlreadyReady = true
				r.RoomEvents.MatchReady <- true
				return
			}
		}
		// todo: 多人游戏: 通用接口
	}
}

// 当 WaitMatchReady 完成, 传入的客户端将会收到匹配完成数据包
func (r *PlaneFighterRoom) WaitMatchReady(cli *clients.ArcadeClient) {
	if r.RoomEvents.MatchAlreadyReady {
		// SHOULD never happened?
		cli.WritePacket(&packets_arcade.ArcadeMatchEvent{Action: packets_arcade.ArcadeMatchEventReady})
		return
	}
	<-r.RoomEvents.MatchReady
	cli.WritePacket(&packets_arcade.ArcadeMatchEvent{Action: packets_arcade.ArcadeMatchEventReady})
}

func (r *PlaneFighterRoom) AddGameReadyCount() {
	r.RoomEvents.GameReadyCount++
	if r.RoomEvents.GameReadyCount >= r.MaxPlayer && !r.RoomEvents.GameAlreadyReady {
		r.RoomEvents.GameReady <- true
		r.RoomEvents.GameAlreadyReady = true
		// OMG, THIS IS SOOOOOO MESSY!!!
		// todo: 以更优雅的方式启动 RoomEntry
		go RoomEntry(r)
	}
}

func (r *PlaneFighterRoom) WaitGameReady() {
	if r.RoomEvents.GameAlreadyReady {
		return
	}
	// println("Waiting")
	<-r.RoomEvents.GameReady
}

func (r *PlaneFighterRoom) MakePlayerList() *packets_arcade.PlaneFighterPlayerList {
	entries := make([]arcade_types.PlaneFighterPlayerEntry, len(r.Clients))
	for i, cli := range r.Clients {
		runtimeId := r.GetClientRuntimeID(cli)
		entries[i] = arcade_types.PlaneFighterPlayerEntry{NickName: cli.Nickname(), UID: cli.UID(), RuntimeID: runtimeId}
	}
	return &packets_arcade.PlaneFighterPlayerList{Entries: entries}
}

func (r *PlaneFighterRoom) GetClientRuntimeID(cli *clients.ArcadeClient) int32 {
	return r.ClientUID2RuntimeID[cli.UID()]
}

// 同时自动分配 RuntimeID
func (r *PlaneFighterRoom) AddClient(cli *clients.ArcadeClient) {
	r.Clients = append(r.Clients, cli)
	r.ClientUID2RuntimeID[cli.UID()] = r.Stage.NewRuntimeID()
	r.RoomEvents.JoinEvent <- cli
}

func (r *PlaneFighterRoom) RemoveClient(cli *clients.ArcadeClient) {
	var index int
	for i, c := range r.Clients {
		if c == cli {
			index = i
			break
		}
	}
	r.Clients = append(r.Clients[:index], r.Clients[index+1:]...)
	delete(r.ClientUID2RuntimeID, cli.UID())
	r.RoomEvents.JoinEvent <- cli
}

func (r *PlaneFighterRoom) IsFull() bool {
	return len(r.Clients) >= r.MaxPlayer
}

func (r *PlaneFighterRoom) SendStageEvents() {
	stageEvents := r.Stage.Events
	if len(stageEvents) > 0 {
		r.broadcastEvents(r.Stage.Events)
		r.Stage.Events = r.Stage.Events[:0]
	}
}
func (r *PlaneFighterRoom) SendStageNewActors() {
	newActors := r.Stage.NewActors
	if len(newActors) > 0 {
		r.broadcastNewActors(r.Stage.NewActors)
		r.Stage.NewActors = r.Stage.NewActors[:0]
	}
}

func (r *PlaneFighterRoom) SendAddScore() {
	if len(r.Stage.AddScoreEvts) > 0 {
		r.broadcastPacket(&packets_arcade.PlaneFighterScores{
			Scores: r.Stage.AddScoreEvts,
		})
		r.Stage.AddScoreEvts = r.Stage.AddScoreEvts[:0]
	}
}

func (r *PlaneFighterRoom) SendStage() {
	players := r.Stage.PlayerPlanes
	entities := r.Stage.Entities
	nplayers := []arcade_types.PFStageEntity{}
	nentities := []arcade_types.PFStageEntity{}
	for _, player := range players {
		if player != nil {
			nplayers = append(nplayers, player.SimpleMarshal())
		}
	}
	for _, entity := range entities {
		if entity != nil && entity.EntityType != define.PlayerBullet && entity.EntityType != define.EnemyBullet {
			nentities = append(nentities, entity.SimpleMarshal())
		}
	}
	r.broadcastPacket(&packets_arcade.PlaneFighterStage{
		Players:  nplayers,
		Entities: nentities,
	})
}

func (r *PlaneFighterRoom) SendStatuses() {
	var statuses []arcade_types.PFPlayerStatus
	for _, player := range r.Stage.PlayerPlanes {
		status := arcade_types.PFPlayerStatus{
			RuntimeID: player.RuntimeID,
			HP:        player.HP,
			Bullets:   player.Bullet,
		}
		statuses = append(statuses, status)
	}
	r.broadcastPacket(&packets_arcade.PlaneFighterPlayerStatuses{Statuses: statuses})
}

func (r *PlaneFighterRoom) broadcastPacket(pk packets.ServerPacket) {
	for _, c := range r.Clients {
		c.WritePacket(pk)
	}
}

func (r *PlaneFighterRoom) broadcastEvents(evts []Event) {
	r.broadcastPacket(&packets_arcade.PlaneFighterActorEvent{
		Events: evts,
	})

}

func (r *PlaneFighterRoom) broadcastNewActors(actors []MovedEntity) {
	pk_actors := make([]arcade_types.PlaneFighterActor, len(actors))
	for i, actor := range actors {
		pk_actors[i] = actor.Marshal()
	}
	r.broadcastPacket(&packets_arcade.PlaneFighterAddActor{
		Actors: pk_actors,
	})
}

func (r *PlaneFighterRoom) CheckAllOnline() bool {
	for _, cli := range r.Clients {
		if !cli.Online {
			return false
		}
	}
	return true
}
