// @Author Gopher
// @Date 2025/1/28 21:30:00
// @Desc
package main

import (
	"bufio"
	"fmt"
	"net"
)

func handelConnection(conn net.Conn) {

	defer conn.Close()

	reader := bufio.NewReader(conn)

	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Connection Close")
			return
		}
		fmt.Printf("Received: %s", message)
		conn.Write([]byte("Echo: " + message))
	}
}

func main() {

	listener, err := net.Listen("tcp", ":8081")
	if err != nil {
		fmt.Println("Error starting TCP server:", err)
		return
	}

	defer listener.Close()

	fmt.Println("TCP server listening on port 8081")
	for {
		conn, err := listener.Accept()

		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		go handelConnection(conn)
	}
}
