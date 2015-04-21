package main

import (
	"fmt"
	"math"
)

const delta = 1e-6

func Sqrt(x float64) float64 {

	z := x
	oldZ := 0.0

	for math.Abs(oldZ-z) > delta {
		// assinge z to oldZ, calcualte a new z
		oldZ, z = z, z-(math.Pow(z, 2)-x)/(2*z)
	}

	return z

}

func main() {
	const x = 2
	mine, theirs := Sqrt(x), math.Sqrt(x)
	fmt.Println(mine, theirs, mine-theirs)
}
