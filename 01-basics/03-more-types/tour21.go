package main

import "fmt"

// we are returning a closure
func adder() func(int) int {
	sum := 0

	// this is a closure
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}
