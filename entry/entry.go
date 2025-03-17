package entry

import (
	"MineArcade-backend/arcade"

	"github.com/pterm/pterm"
)

func Entry() {
	pterm.Info.Println("MineArcade-backend 启动中...")
	CreateDataDirs()
	arcade.Launch()
	defer arcade.Exit()
	TestServer()
}
