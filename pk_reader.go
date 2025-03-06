package main

import (
	"net"
)

func DialPk(con net.Conn) {
	var buf []byte
	con.Read(buf)
}
