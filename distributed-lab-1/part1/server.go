package main

import (
	"net"
	"fmt"
	"bufio"
	)

func handleConnection(conn *net.Conn) {
	reader := bufio.NewReader(*conn)
	for {
		msg, _ := reader.ReadString('\n')
		fmt.Printf(msg)
		fmt.Fprintln(*conn, "OK")
	}

}

func main() {
	ln, _ := net.Listen("tcp", ":8080")
	for {
		conn, _ := ln.Accept()
		go handleConnection(&conn)
	}
}
