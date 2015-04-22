package main

import (
	"fmt"
)

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}

	// this closes the channel
	close(c)
}

func main() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)

	// we can test to see if the channel is closed
	// the second value (ok) would be false if the channel is closed
	// v, ok := <-c

	// This reads from the channel and will stop when the channel is closed.
	for i := range c {
		fmt.Println(i)
	}
}
