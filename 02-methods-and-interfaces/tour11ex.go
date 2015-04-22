package main

import "code.google.com/p/go-tour/reader"

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (r MyReader) Read(b []byte) (int, error) {
	size := len(b)
	for i, _ := range b {
		b[i] = 'A'
	}
	return size, nil
}

func main() {
	reader.Validate(MyReader{})
}
