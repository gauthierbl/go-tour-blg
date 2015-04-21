package main

import "code.google.com/p/go-tour/pic"

func Pic(dx, dy int) [][]uint8 {
	p := make([][]uint8, dy)
	for i := range p {
		p[i] = make([]uint8, dx)
	}

	for y, row := range p {
		for x := range row {
			row[x] = uint8(x * y)
		}
	}

	return p
}

func main() {
	// this is going to dump a base64 png onto the console, this code is best run
	// in the online go console: https://tour.golang.org/moretypes/14
	pic.Show(Pic)
}
