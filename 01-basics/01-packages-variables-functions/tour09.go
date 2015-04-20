package main

import "fmt"

var i, j int = 1, 2

func main() {
	// we can omit type info when setting an initializer
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}
