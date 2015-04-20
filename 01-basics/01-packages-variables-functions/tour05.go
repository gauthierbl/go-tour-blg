package main

import "fmt"

// we can omit the type for all but the last, this is the same as in tour04
func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}
