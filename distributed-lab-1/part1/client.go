package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
)

func read(conn *net.Conn) {
	reader := bufio.NewReader(*conn)
	msg, _ := reader.ReadString('\n')
	fmt.Printf(msg)
}

func main() {
	msgPtr := flag.String("msg", "No message supplied", "Message to send")
	flag.Parse()
	conn, _ := net.Dial("tcp", "127.0.0.1:8080")
	fmt.Fprintln(conn, *msgPtr)
	read(&conn)
}