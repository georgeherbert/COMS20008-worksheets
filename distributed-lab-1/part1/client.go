package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func read(conn *net.Conn) {
	reader := bufio.NewReader(*conn)
	msg, _ := reader.ReadString('\n')
	fmt.Printf(msg)
}

func main() {
	//msgPtr := flag.String("msg", "No message supplied", "Message to send")
	//flag.Parse()
	stdin := bufio.NewReader(os.Stdin)
	conn, _ := net.Dial("tcp", "127.0.0.1:8080")
	for {
		fmt.Printf("Enter text: ")
		text, _ := stdin.ReadString('\n')
		fmt.Fprintf(conn, text)
		read(&conn)
	}
}