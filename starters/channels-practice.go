package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	n := 5

	// Only n messages can be in a channel
	// at any time
	ch := make(chan int, n)
	wg.Add(2)

	// Data recieved from channel
	go func(ch <-chan int) {
		for {
			if i, ok := <-ch; ok {
				fmt.Println(i)
			} else {
				break
			}
		}
		wg.Done()
	}(ch)

	// Data sent through channel
	go func(ch chan<- int) {
		i := 42
		ch <- i
		ch <- 5
		close(ch)
		wg.Done()
	}(ch)

	wg.Wait()
}
