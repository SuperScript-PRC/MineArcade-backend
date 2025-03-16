package entry

import (
	"MineArcade-backend/arcade"
	"MineArcade-backend/server"

	"github.com/pterm/pterm"
)

func Entry() {
	pterm.Info.Println("MineArcade-backend 启动中...")
	CreateDataDirs()
	arcade.Launch()
	server.TestServer()
}
