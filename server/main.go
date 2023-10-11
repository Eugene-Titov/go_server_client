package main

import (
	"fmt"
	"net"
	"os"
)

func handleCommand(command string, conn net.Conn) {
	fmt.Println(command)
	conn.Write([]byte("bad command"))
}

func serveConnection(conn net.Conn) {
	defer conn.Close()
	fmt.Println("new connection")
	
	for {
		var bufResult []byte
		for {
			buf := make([]byte, 1024)
			n, _ := conn.Read(buf)
			if n > 0 {
				bufResult = append(bufResult, buf[:n]...)
				if n < 1024 {
					break
				}
			} else {
				break
			}
		}
		handleCommand(string(bufResult), conn)
	}
	fmt.Println("exit new connection")
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