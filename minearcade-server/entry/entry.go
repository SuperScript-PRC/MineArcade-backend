package entry

import (
	"MineArcade-backend/minearcade-server/core"
	"log/slog"
	"net"
)

func Entry() {
	slog.Info("MineArcade-backend 启动中...")
	CreateDataDirs()
	mcore := core.NewCore()
	clientEntry := func(tcp_conn net.Conn, udp_conn *net.UDPConn) {
		ClientConnEntry(mcore.Clients, tcp_conn, udp_conn)
	}
	mcore.Server.SetConnHandler(clientEntry)
	mcore.Run()
	WaitClosed()
}
