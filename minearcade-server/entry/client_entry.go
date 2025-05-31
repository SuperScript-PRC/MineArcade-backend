package entry

import (
	"MineArcade-backend/minearcade-server/clients"
	"MineArcade-backend/minearcade-server/minearcade"
	"net"
)

func ClientConnEntry(clients_group *clients.ArcadeClients, tcp_conn net.Conn, udp_conn *net.UDPConn) {
	cli := clients.MakeClient(tcp_conn, udp_conn)
	clients_group.AddClient(cli)
	defer clients_group.RemoveClient(cli)
	ok := clients.ClientLogin(cli)
	if !ok {
		return
	}
	// ClientLoop
	minearcade.MainLobbyEntry(cli)
}
