package main

import (
	"fmt"
	"os"
)

func main() {
	defer fmt.Println("This will NOT run when we use os.Exit()!")

	fmt.Println("This program will exit immediately with status 3.")
	fmt.Println()
	fmt.Println("If we use `go run ...`, the exit status will be printed by go's runner.")
	fmt.Println("If we use `go build` and run this as a binary, the exit status won't be printed.")

	os.Exit(3)
	// We have to exit like this if we want to exit with
	// a status other than the default 1.
}
