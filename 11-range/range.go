package main

import "fmt"

func main() {
	nums := []int{10, 20, 30, 40}
	fmt.Println("slice of integers, nums:", nums)
	var sum int
	for _, v := range nums {
		sum += v
	}
	fmt.Println("iterate over them using range and sum:", sum)

	for i, v := range nums {
		if v == 30 {
			fmt.Println("value 30 is in index:", i)
		}
	}
	fmt.Println()

	m := map[string]string{"a": "apple", "b": "banana"}
	fmt.Println("iterate over map m:", m)
	for k, v := range m {
		fmt.Printf(" %s -> %s\n", k, v)
	}
	fmt.Println()

	s := "go"
	fmt.Println("iterate over string s:", s)
	for i, c := range s {
		fmt.Printf(" %d: %c\n", i, c)
	}
}
