package core

import (
	"MineArcade-backend/minearcade-server/arcade"
	"MineArcade-backend/minearcade-server/clients"
	"MineArcade-backend/minearcade-server/server"
)

type MineArcadeCore struct {
	Server  *server.MineArcadeServer
	Clients *clients.ArcadeClients
}

func NewCore() *MineArcadeCore {
	server := server.NewServer()
	clients := clients.NewClients()
	server.SetClientUDPPacketHandler(clients.HandoutUDPBytePacket)
	return &MineArcadeCore{
		Server:  server,
		Clients: clients,
	}
}

func (c *MineArcadeCore) Run() {
	c.Server.StartServer()
	arcade.Launch()
}
