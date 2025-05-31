package arcade

import "MineArcade-backend/minearcade-server/arcade/public_minearea"

// 初始化并激活所有子游戏。
func Launch() {
	go public_minearea.Launch()
}
