package main

import "fmt"

func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		// push a statement on to the "defer stack", defers will be executed in last in first out
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
