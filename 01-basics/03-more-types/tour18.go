package main

import "fmt"

func main() {
	m := make(map[string]int)

	//insert or update a element in a map
	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	// get an element
	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	// remove an element
	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	// get and tests to see if the value exists
	m["blg"] = 316
	v, ok := m["blg"]
	fmt.Println("blg: The value:", v, "Present?", ok)

	// get but in this case it does not exist in the map,
	// will get the "zero value" for the type and false
	v, ok = m["Answer"]
	fmt.Println("Answer: The value:", v, "Present?", ok)
}
