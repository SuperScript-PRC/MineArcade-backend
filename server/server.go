package server

import (
	"MineArcade-backend/clients"
	"fmt"
	"net"

	"github.com/pterm/pterm"
)

func TestServer() {
	listener, err := net.Listen("tcp", ":6000")
	if err != nil {
		panic(err)
	}
	for {
		con, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			break
		}
		pterm.Info.Println("新连接:", con.RemoteAddr().String())
		go clients.HandleConnection(con)
	}
	fmt.Println("Exit.")
}
