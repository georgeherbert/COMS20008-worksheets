package main

import (
	"fmt"
	"time"
)

// slowSender sends a string every 2 seconds.
func slowSender(c chan<- string) {
	for {
		time.Sleep(2 * time.Second)
		c <- "I am the slowSender"
	}
}

// fastSender sends consecutive ints every 500 ms.
func fastSender(c chan<- int) {
	for i := 0; ; i++ {
		time.Sleep(500 * time.Millisecond)
		c <- i
	}
}

func fasterSender(c chan<- []int) {
	for {
		time.Sleep(200 * time.Millisecond)
		c <- []int{1, 2, 3}
	}
}

// main starts the two senders and then goes into an infinite loop of receiving their messages.
func main() {
	ints := make(chan int)
	go fastSender(ints)
	strings := make(chan string)
	go slowSender(strings)
	slices := make(chan []int)
	go fasterSender(slices)

	for { // = while(true)
		select {
		case st := <-strings:
			fmt.Println("Received a string", st)
		case i := <-ints:
			fmt.Println("Received an int", i)
		case sl := <-slices:
			fmt.Println("Received a slice", sl)
		default:
			fmt.Println("--- Nothing to receive, sleeping for 3s...")
			time.Sleep(3 * time.Second)
		}
	}
}
