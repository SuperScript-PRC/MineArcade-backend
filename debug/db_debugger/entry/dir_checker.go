package entry

import (
	"MineArcade-backend/minearcade-server/defines"
	"fmt"
	"os"
)

func CheckDir() bool {
	stat, err := os.Stat(defines.BASE_PATH)
	if err != nil {
		fmt.Println("请把该程序置于 MineArcade 后端工作区目录下")
		return false
	}
	if !stat.IsDir() {
		fmt.Println("请把该程序置于 MineArcade 后端工作区目录下")
		return false
	}
	return true
}
