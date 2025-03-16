package arcade

import "MineArcade-backend/arcade/public_minearea"

func Launch() {
	go public_minearea.Launch()
}
