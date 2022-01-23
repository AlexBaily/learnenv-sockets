package sockets

import (
	"net"
	"fmt"
)

func StartServer(connPort string) {

	ln, _ := net.Listen("tcp", connPort)


	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
		}
	
		go connHandler(conn)
	}
}

func connHandler(conn net.Conn) {
	defer conn.Close()

	conn.Write([]byte("Connection successful."))
}