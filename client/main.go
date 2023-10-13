package main

import (
	"fmt"
	"net"
	"os"
	"io"
//	"bufio"
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
		
		var output string
		for {
			buff := make([]byte, 1024)
			n, _ := conn.Read(buff)
			fmt.Println(n)
			if n < 1024 {
				output += string(buff[0:n])
				break
			} else {
				output += string(buff)
			}
		}
		fmt.Println(output)
	}
	fmt.Println("exit")
}