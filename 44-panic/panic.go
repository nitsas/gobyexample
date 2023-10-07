package main

import (
	"fmt"
	"os"
)

func main() {
	panic("I panicked!")

	// unreachable code
	filename := "/tmp/gobyexample-tmp-file.txt"
	_, err := os.Create(filename)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Created tempfile", filename)
	}
}
