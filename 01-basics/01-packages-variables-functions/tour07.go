package main

import "fmt"

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x

	// a "naked" return will return x and y above
	return
}

func main() {
	fmt.Println(split(17))
}
