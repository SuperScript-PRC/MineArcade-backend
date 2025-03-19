package main

import (
	terminal_entry "MineArcade-backend/debug/db_debugger/entry"
)

func main() {
	if !terminal_entry.CheckDir() {
		return
	}
	terminal_entry.Main()
}
