package main

import "fmt"

func main() {
	// this defer statement will be called when the function returns
	defer fmt.Println("world")

	fmt.Println("hello")
}
