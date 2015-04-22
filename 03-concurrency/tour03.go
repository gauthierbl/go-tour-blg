package main

import "fmt"

func main() {

	// make a buffered channel of type int with buffer size == 2
	c := make(chan int, 2)
	c <- 1
	c <- 2

	// Sends to a buffered channel block only when the buffer is full. Receives block when the buffer is empty.
	// We can overflow the channel and cause a deadlock (which go detects)
	// c <- 3

	fmt.Println(<-c)
	fmt.Println(<-c)
}
