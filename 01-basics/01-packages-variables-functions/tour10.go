package main

import "fmt"

func main() {
	var i, j int = 1, 2

	// don't need var in a function with := the short assignment statement
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}
