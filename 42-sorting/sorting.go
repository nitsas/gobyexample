package main

import (
	"fmt"
	"slices"
)

func main() {
	strs := []string{"c", "a", "b"}
	fmt.Println("strs:", strs)
	fmt.Println("slices.Sort(strs)")
	slices.Sort(strs)
	fmt.Println("strs:", strs)
	fmt.Println()

	ints := []int{3, 1, 4, 2}
	fmt.Println("ints:", ints)
	fmt.Println("slices.Sort(ints)")
	slices.Sort(ints)
	fmt.Println("ints:", ints)
	fmt.Println()

	fmt.Println("slices.IsSorted(ints):", slices.IsSorted(ints))
}
