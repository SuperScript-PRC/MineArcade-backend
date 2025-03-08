package main

import (
	"MineArcade-backend/server"

	"github.com/pterm/pterm"
)

func main() {
	pterm.Info.Println("服务端已启动.")
	server.TestServer()
}
