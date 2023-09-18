package main

import "fmt"

func sum(nums ...int) int {
	var result int
	for _, ni := range nums {
		result += ni
	}
	return result
}

func main() {
	fmt.Println("sum(1, 2):", sum(1, 2))
	fmt.Println("sum(3, 4, 5):", sum(3, 4, 5))

	nums := []int{6, 7, 8, 9}
	fmt.Println("sum(", nums, "...):", sum(nums...))
}
