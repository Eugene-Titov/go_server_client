package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8100")
	if err != nil {
		fmt.Println(err)
		fmt.Println("error to create client")
		os.Exit(1)
	}
	defer conn.Close()
	for {
		var command string
		fmt.Print("enter command: ")
		fmt.Scan(&command)
		conn.Write([]byte(command))
	}
	fmt.Println("exit")
}