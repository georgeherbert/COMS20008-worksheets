package main

import (
	"net"
	"fmt"
	"bufio"
	)

func main() {
	ln, _ := net.Listen("tcp", ":8080")
	conn, _ := ln.Accept()
	reader := bufio.NewReader(conn)
	msg, _ := reader.ReadString('\n')
	fmt.Println(msg)
}
