package main

import (
	"fmt"
	"os"
)

// If we run this with `go run ...`, the executable's name is a tempfile.
// Better use `go build ...` to make it.
func main() {
	fmt.Printf("len(os.Args): %d\n", len(os.Args))
	fmt.Printf("os.Args: %#v\n", os.Args)

	// If we try to access an argument beyond len(os.Args), we'll
	// get an error: "index out of bounds".
	// fmt.Printf("os.Args[2]: %s\n", os.Args[2])
}
