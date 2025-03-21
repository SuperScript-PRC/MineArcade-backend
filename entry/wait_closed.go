package entry

import (
	"MineArcade-backend/arcade"
	"os"
	"os/signal"
	"syscall"

	"github.com/pterm/pterm"
)

func WaitClosed() os.Signal {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	sig := <-signalChan
	arcade.Exit()
	pterm.Success.Println("MineArcade-backend 已退出")
	return sig
}
