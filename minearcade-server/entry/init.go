package entry

import (
	"MineArcade-backend/minearcade-server/defines"
	"os"
)

func CreateDataDirs() {
	os.MkdirAll(defines.ACCOUNT_DIR_PATH, 0777)
	os.MkdirAll(defines.PLAYER_STORE_DIR_PATH, 0777)
	os.MkdirAll(defines.ARCADE_DATA_PATH, 0777)
}
