package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func main() {
	path := filepath.Join("gobyexample", "63-file-paths", "file-paths.go")
	fmt.Printf("path: %#v\n", path)
	fmt.Printf("path type: %T\n", path)
	// path is a string

	dir, filename := filepath.Split(path)
	// There's also filepath.Dir() and filepath.Base() that get each part.
	fmt.Printf("dir: %#v, filename: %#v\n", dir, filename)
	ext := filepath.Ext(filename)
	filenameWithoutExt := strings.TrimSuffix(filename, ext)
	fmt.Printf("ext: %#v, filenameWithoutExt: %#v\n", ext, filenameWithoutExt)

	fmt.Println("IsAbs(path):", filepath.IsAbs(path))
	fmt.Println()

	otherPath := filepath.Join("gobyexample", "62-line-filters")
	relPath, err := filepath.Rel(path, otherPath)
	if err != nil {
		panic(err)
	}
	fmt.Printf("How to go to %#v: %#v\n", otherPath, relPath)
	fmt.Println()

	unrelatedPath := filepath.Join("/", "unrelated-dir", "file.txt")
	relPath, err = filepath.Rel(path, unrelatedPath)
	if err != nil {
		fmt.Printf("Cannot go from path to %#v!\n", unrelatedPath)
		fmt.Printf("Error: %e\n", err)
	} else {
		fmt.Printf("How to go to %#v: %#v\n", unrelatedPath, relPath)
	}
}
