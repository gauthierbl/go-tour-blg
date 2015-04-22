package main

import (
	"fmt"
	"math"
)

//go does not have classes but methods can be defined for structs
type Vertex struct {
	X, Y float64
}

// this is a method on the Vertex struct type, the (v *Vertex) part is the "method receiver"
func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := &Vertex{3, 4}
	fmt.Println(v.Abs())
}
