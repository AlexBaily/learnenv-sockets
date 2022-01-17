package sockets

import (
	"net"
	"fmt"
)

func StartServer(connPort string) {

	ln, _ := net.Listen("tcp", connPort)

	conn, _ := ln.Accept()

	for {
		fmt.Println(conn)
	}
}

func connHandler(conn) {
	
}