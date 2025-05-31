package tests

import (
	"MineArcade-backend/minearcade-server/protocol"
	"fmt"
	"net"
)

func HandleConnection(conn net.Conn) {
	defer func() {
		e := recover()
		if e != nil {
			fmt.Println(e)
		}
		_ = conn.Close()
	}()
	// writer := protocol.Writer{}
	// writer.String("Hello 大家！")
	// conn.Write(writer.GetFullBytes())
	reader := protocol.Reader{}
	var str string
	bs := make([]byte, 1024)
	n, _ := conn.Read(bs)
	reader.SetFullBytes(bs, n)
	fmt.Printf("bts=%v\n", bs)
	reader.StringUTF(&str)
	fmt.Println("str=" + str)
	writer := protocol.Writer{}
	writer.StringUTF("Hello 大家！")
	conn.Write(writer.GetFullBytes())
	err := conn.Close()
	if err != nil {
		fmt.Println("Error closing connection")
		panic(err)
	}
}

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
		go HandleConnection(con)
	}
	fmt.Println("Exit.")
}
