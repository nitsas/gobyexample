package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func must(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// Create a tempfile. First argument "" means put it in the default place.
	f, err := os.CreateTemp("", "gobyexample-tfad")
	must(err)
	defer os.Remove(f.Name())
	fmt.Println("Created tempfile:", f.Name())

	numBytes, err := f.Write([]byte{1, 2, 3, 4})
	must(err)
	fmt.Println("Wrote", numBytes, "bytes to tempfile")

	dname, err := os.MkdirTemp("", "gobyexample-tfad-dir")
	must(err)
	fmt.Println("Created tempdir:", dname)
	defer os.RemoveAll(dname)

	f2name := filepath.Join(dname, "file2")
	must(os.WriteFile(f2name, []byte("Hello world!"), 0666))
	fmt.Println("Wrote \"Hello World\" to", f2name)
	defer os.Remove(f2name)
}
