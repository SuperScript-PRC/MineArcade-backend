package plane_fighter

import (
	"fmt"

	"github.com/google/uuid"
)

var GlobalRooms = make(map[string]*PlaneFighterRoom)

func GetAvailRoom() (string, *PlaneFighterRoom) {
	for roomName, room := range GlobalRooms {
		if !room.IsFull() {
			return roomName, room
		}
	}
	// all full
	ud, err := uuid.NewUUID()
	if err != nil {
		panic(fmt.Errorf("RoomID generate error: %v", err))
	}
	// TODO: UUID may be duplicated
	newRoom := NewRoom(ud.String(), 2)
	return newRoom.RoomID, newRoom
}
