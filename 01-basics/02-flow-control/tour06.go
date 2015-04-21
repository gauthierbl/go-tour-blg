package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {

	// We are assigning v then doing hte boolean, like a for
	if v := math.Pow(x, n); v < lim {
		// v is only in scoper here
		return v
	}

	// v is not in scope here
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
