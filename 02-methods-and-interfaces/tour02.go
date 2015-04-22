package main

import (
	"fmt"
	"math"
)

// in this example we delcare a method on our non-struct type
// we cannot define a method from a different package or built in types
type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}
