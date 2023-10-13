package main

import (
	"fmt"
	"net"
	"os"
	"bufio"
	"strings"
)

func handleCommand(in string, conn net.Conn) {
	if len(in) == 0 {
		return
	}
	
	in = strings.TrimSpace(in)
	commands := strings.Split(in, " ")
	var output string
	switch commands[0] {
		case "ls":
			output = getFiles()
		case "cd":
			if len(commands) > 1 {
				output = cd(commands[1])
			} else {
				output = "error path"
			}
		case "get":
			if len(commands) > 1 {
				var b []byte
				output, b = getFile(commands[1])
				conn.Write([]byte("create file: " + commands[1]))
				conn.Write(b)
			} else {
				output = "error to copy file"
			}
		default:
			output = "bad command\n"
	}
	
	type WS interface {
		WriteString(s string) (n int, err error)
	}
	
	if ws, ok := conn.(WS); ok {
		ws.WriteString(output)
	} else {
		conn.Write([]byte(output))
	}
}

func serveConnection(conn net.Conn) {
	defer conn.Close()
	fmt.Println("new connection")
	
	for {
		command, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			break
		}
		handleCommand(command, conn)
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

func cd(path string) string {
	err := os.Chdir(path)
	if err != nil {
		return "error path"
	}
	return "ok\n"
}

func getFiles() (res string) {
	files, err := os.ReadDir(".")
	if err != nil {
		res = "error"
		return
	}
	for _, f := range files {
		res += f.Name() + "\n"
	}
	return
}

func getFile(path string) (string, []byte) {
	var bytes []byte
	bytes, err := os.ReadFile(path)
	if err != nil {
		return "error file\n", bytes
	}
	return "\nok\n", bytes
}