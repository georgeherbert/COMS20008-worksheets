package main

import (
	"fmt"
	"net"
	"flag"
	)

func main() {
	msgPtr := flag.String("msg", "No message supplied", "Message to send")
	flag.Parse()
	conn, _ := net.Dial("tcp", "127.0.0.1:8080")
	fmt.Fprintf(conn, *msgPtr)
}