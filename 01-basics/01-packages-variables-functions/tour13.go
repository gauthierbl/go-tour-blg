package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4

	// Go assignments are explkicited typed and require a converersion.
	// Omitting T(v) would cause a comple error.
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z int = int(f)

	// more simply
	b := int(f)

	fmt.Println(x, y, z, b)
}
