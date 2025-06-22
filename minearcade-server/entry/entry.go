package entry

import (
	"MineArcade-backend/minearcade-server/core"
	"log/slog"
	"net"
)

func Entry() {
	slog.Info("MineArcade-backend 启动中...")
	CreateDataDirs()
	arcadeCore := core.NewCore()
	clientEntry := func(tcp_conn net.Conn, udp_conn *net.UDPConn) {
		ClientConnEntry(arcadeCore.Clients, tcp_conn, udp_conn)
	}
	arcadeCore.Server.SetConnHandler(clientEntry)
	arcadeCore.Run()
	<-WaitClosed()
	CloseAll()
}
