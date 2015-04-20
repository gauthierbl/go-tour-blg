package main

import "fmt"

func main() {

	// type inference with :=
	i := 42 // change me!
	fmt.Printf("i is of type %T\n", i)

	s := "42" // change me!
	fmt.Printf("s is of type %T\n", s)

	f := float64(42) // change me!
	fmt.Printf("f is of type %T\n", f)

	b := true // change me!
	fmt.Printf("b is of type %T\n", b)
}
