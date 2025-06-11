package entry

import (
	"MineArcade-backend/minearcade-server/arcade"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

func WaitClosed() os.Signal {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	sig := <-signalChan
	arcade.Exit()
	slog.Info("MineArcade-backend 已退出")
	return sig
}
