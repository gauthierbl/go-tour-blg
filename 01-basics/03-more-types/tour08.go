package main

import "fmt"

func main() {
	p := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("p ==", p)
	fmt.Println("p[1:4] ==", p[1:4]) // a slice that is elements 1-4

	// missing low index implies 0
	fmt.Println("p[:3] ==", p[:3]) // a slice that is elements 0-3

	// missing high index implies len(s)
	fmt.Println("p[4:] ==", p[4:]) // a slice that is elements 4-tail

	fmt.Println("p[4:4] ==", p[4:4]) // a empty slice
	fmt.Println("p[4:5] ==", p[4:5]) // a one element slice
}
