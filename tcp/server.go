package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func main() {
	lis, err := net.Listen("tcp", ":8080")
	defer lis.Close()
	if err != nil {
		panic(err)
	}

	for {
		conn, err := lis.Accept()
		fmt.Println(conn.RemoteAddr().String(), " Connected")
		if err != nil {
			panic(err)
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	for {
		dataBytes, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			panic(err)
		}
		data := strings.TrimSpace(string(dataBytes))

		fmt.Println(data)
		fmt.Println(len(data))
		if data == "exit" {
			conn.Close()
			fmt.Println(conn.RemoteAddr().String(), " Disconnected")
			break
		}

	}
}
