package main

import "fmt"

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{10, 20} // has type Vertex
	v2 = Vertex{X: 3}   // Y:0 is implicit
	v3 = Vertex{}       // X:0 and Y:0
	p  = &Vertex{4, 5}  // has type *Vertex
)

func main() {
	fmt.Println(v1, v2, v3, p)
}
