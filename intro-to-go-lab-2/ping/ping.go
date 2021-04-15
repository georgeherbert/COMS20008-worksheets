package main

import (
	"fmt"
	"log"
	"os"
	"runtime/trace"
	"time"
)

func foo(channel chan string) {
	// TODO: Write an infinite loop of sending "pings" and receiving "pongs"
	for i := 0; i >= 0; i++ {
		sendMessage := "ping"
		fmt.Println("Foo is sending:", sendMessage)
		channel <- sendMessage

		recieveMessage := <-channel
		fmt.Println("Foo has received:", recieveMessage)

		fmt.Println("")
	}
}

func bar(channel chan string) {
	for i := 0; i >= 0; i++ {
		receiveMessage := <-channel
		fmt.Println("Bar has received:", receiveMessage)

		sendMessage := "pong"
		fmt.Println("Bar is sending:", sendMessage)
		channel <- sendMessage
	}

	// TODO: Write an infinite loop of receiving "pings" and sending "pongs"
}

func pingPong() {
	// TODO: make channel of type string and pass it to foo and bar

	pingPongChannel := make(chan string)

	go foo(pingPongChannel) // Nil is similar to null. Sending or receiving from a nil chan blocks forever.
	go bar(pingPongChannel)

	time.Sleep(500 * time.Millisecond)
}

func main() {
	f, err := os.Create("trace.out")
	if err != nil {
		log.Fatalf("failed to create trace output file: %v", err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatalf("failed to close trace file: %v", err)
		}
	}()

	if err := trace.Start(f); err != nil {
		log.Fatalf("failed to start trace: %v", err)
	}
	defer trace.Stop()

	pingPong()
}
