package main

import (
	"fmt"
	"slices"
)

func main() {
	var s []string
	fmt.Println("empty slice of strings, s:", s)
	fmt.Println("s == nil:", s == nil)
	fmt.Println("len(s):", len(s))
	fmt.Println()

	s = make([]string, 3, 10)
	fmt.Println("make a slice of strings of length 3 and capacity 10, s:", s)
	fmt.Println("s == nil:", s == nil)
	fmt.Println("len(s):", len(s))
	fmt.Println("cap(s):", cap(s))
	fmt.Println()

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set values for s:", s)
	fmt.Println("s[2] is", s[2])
	fmt.Println("len(s):", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	// Note that we had to get the return value from append.
	// It might be a different object than before.
	fmt.Println("appended to s:", s)

	sCopy := make([]string, len(s))
	copy(sCopy, s)
	fmt.Println("copy of s, sCopy:", sCopy)
	fmt.Println()

	fmt.Println("s[2:5]:", s[2:5])
	fmt.Println("s[:5]:", s[:5])
	fmt.Println("s[2:]:", s[2:])
	fmt.Println()

	t := []string{"g", "h", "i"}
	fmt.Println("prefilled slice, t:", t)

	t2 := []string{"g", "h", "i"}
	fmt.Println("prefilled slice, t2:", t2)
	fmt.Println("slices.Equal(t, t2):", slices.Equal(t, t2))
	fmt.Println()

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("two-dimensional slice with variable inner slice length, twoD:", twoD)

	// Take a look at this guide for Go slices usage and internals:
	// https://go.dev/blog/slices-intro
}
