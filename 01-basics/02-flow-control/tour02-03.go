package main

import "fmt"

func main() {
	sum := 1

	// it is the same to write the following but gofmt will fix it to below:
	// for ; sum < 1000; {
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
