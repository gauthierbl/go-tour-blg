package main

import "fmt"

func main() {

	// an array's lenght is part of it's type so arrays cannot be resized
	var a [2]string // 2 element array of string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)
}
