package entry

import (
	"MineArcade-backend/minearcade-server/arcade"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

func WaitClosed() chan os.Signal {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	return signalChan
}

func CloseAll() {
	arcade.Exit()
	slog.Info("MineArcade-backend 已退出")
}
