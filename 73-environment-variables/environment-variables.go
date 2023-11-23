package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Printf("os.Environ: %#v\n", os.Environ())
	fmt.Printf("os.Getenv(\"FOO\"): %#v\n", os.Getenv("FOO"))
	fmt.Printf("os.Getenv(\"BAR\"): %#v\n", os.Getenv("BAR"))
	fmt.Println()

	fmt.Println("os.Setenv(\"FOO\", \"lala\")")
	os.Setenv("FOO", "lala")
	fmt.Printf("os.Getenv(\"FOO\"): %#v\n", os.Getenv("FOO"))
	fmt.Println()

	fmt.Println("Loop over os.Environ() and keep FOO & BAR:")
	for _, keyval := range os.Environ() {
		keyvalSplit := strings.SplitN(keyval, "=", 2)
		key, val := keyvalSplit[0], keyvalSplit[1]
		if key == "FOO" || key == "BAR" {
			fmt.Printf("%s: %s\n", key, val)
		}
	}
}
