package sockets

import (
	"net"
	"fmt"
)

func StartServer(connPort string) {

	ln, _ := net.Listen("tcp", connPort)

	conn, err := ln.Accept()
	if err != nil {
		fmt.Println(err)
	}

	for {
		fmt.Println(conn)
	}
}

//func connHandler(conn) {

//}