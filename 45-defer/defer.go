package main

import (
	"fmt"
	"os"
)

func createFile(filename string) *os.File {
	fmt.Printf("createFile(%s)\n", filename)
	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	return f
}

func closeFile(f *os.File) {
	fmt.Printf("closeFile(%p)\n", f)
	err := f.Close()
	if err != nil {
		// What if the deferred closeFile panics instead of using os.Exit(1)?
		// Answer: It will also print a backtrace instead of just exiting.
		// panic("closeFile panicked!!!")

		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func writeFile(f *os.File, contents string) {
	fmt.Printf("writeFile(%p, %s)\n", f, contents)
	fmt.Fprintln(f, contents)
}

func main() {
	filename := "/tmp/gobyexample-tmp-defer.txt"

	// If you want to see an error while creating the file:
	// touch /tmp/gobyexample-tmp-defer.txt
	// chmod u-w /tmp/gobyexample-tmp-defer.txt
	f := createFile(filename)

	// The deferred closeFile(f) will run when main() exits in any way.
	defer closeFile(f)
	writeFile(f, "Hello world!")

	// Will the deferred closeFile(f) run if we get a panic here?
	// Answer: Yes it will.
	// panic("Panic before the end. Will deferred closeFile(f) run?")

	// If you want to see an error in closeFile, uncomment this 2nd close:
	// f.Close()

	// What happens if I defer one more thing?
	// Answer: All the deferred functions will run, in reverse order.
	// So this will run before closeFile().
	// defer func() { fmt.Println("2nd deferred function!") }()
}
