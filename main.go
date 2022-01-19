package main

import (
	"fmt"
	//"net"
	"os"
	"github.com/alexbaily/learnens-sockets/sockets"
)

var (
	connPort string
)

func init() {
	//Grab env variable but set default to 8080
	if connPort = os.Getenv("CONN_PORT"); connPort == "" {
		connPort = ":8080"
	}
}

func main() {
	fmt.Println("Server running on: " + connPort)
	sockets.StartServer(connPort)
}
