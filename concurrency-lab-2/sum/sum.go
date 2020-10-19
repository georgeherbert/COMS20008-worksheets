package main

import (
	"fmt"
	"sync"
)

func increment(value chan int, wg *sync.WaitGroup) {
	n := <-value
	n += 1
	wg.Done()
	value <- n
}

func main() {
	sum := make(chan int)
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go increment(sum, &wg)
	}
	sum <- 0
	wg.Wait()
	finalSum := <-sum
	fmt.Println(finalSum)
}
