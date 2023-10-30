package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	now := time.Now()
	nowFormatted := now.Format(time.RFC3339)

	filename1 := "/tmp/gobyexample-writing-files-1.txt"
	str1 := fmt.Sprintf("Hello world from Go at %s\n", nowFormatted)
	err := os.WriteFile(filename1, []byte(str1), 0644)
	check(err)
	fmt.Println("Wrote to file 1. Check it out like this:")
	fmt.Printf("cat %#v\n", filename1)
	fmt.Println()

	filename2 := "/tmp/gobyexample-writing-files-2.txt"
	str2 := fmt.Sprintf("Hello again, this is Go at %s!\n", nowFormatted)

	f2, err := os.Create(filename2)
	check(err)
	nb, err := f2.Write([]byte(str2))
	check(err)
	fmt.Printf("Wrote %d bytes to file f2\n", nb)
	nb, err = f2.WriteString("Another line of data\n")
	check(err)
	fmt.Printf("Wrote another %d bytes to file f2\n", nb)
	f2.Sync()
	fmt.Println("f2.Sync() to flush the data to disk")
	fmt.Printf("cat %#v\n", filename2)
	fmt.Println()

	wbuf := bufio.NewWriter(f2)
	fmt.Println("wbuf := bufio.NewWriter(f2)")
	nb, err = wbuf.WriteString("Buffered line\n")
	check(err)
	fmt.Printf("Wrote %d bytes to wbuf. wbuf.Buffered(): %d, wbuf.Available(): %d\n", nb, wbuf.Buffered(), wbuf.Available())
	wbuf.Flush()
	fmt.Printf("wbuf.Flush() to flush the data to disk. wbuf.Buffered(): %d, wbuf.Available(): %d\n", wbuf.Buffered(), wbuf.Available())
	fmt.Printf("cat %#v\n", filename2)
	fmt.Println()
}
