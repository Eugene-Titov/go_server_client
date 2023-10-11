package main

import (
	"fmt"
	"net"
	"os"
)

func serveConnection(conn net.Conn) {
	fmt.Println("new connection")
}

func main() {
	ln := createServer("tcp", "localhost:8100")
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("error connection")
			continue
		}
		serveConnection(conn)
	}
	fmt.Println("exit")
}

func createServer(protocol, port string) net.Listener {
	ln, err := net.Listen(protocol, port)
	if err != nil || ln == nil {
		fmt.Println("error to create server")
		os.Exit(1)
	}
	return ln
}