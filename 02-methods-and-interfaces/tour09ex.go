package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("Sqrt: negative number %g", e)
}

const delta = 1e-6

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}

	z := x
	oldZ := 0.0

	for math.Abs(oldZ-z) > delta {
		// assinge z to oldZ, calcualte a new z
		oldZ, z = z, z-(math.Pow(z, 2)-x)/(2*z)
	}

	return z, nil

}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
