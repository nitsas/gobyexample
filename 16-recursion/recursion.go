package main

import "fmt"

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	fmt.Println("fact(3):", fact(3))
	fmt.Println("fact(4):", fact(4))
	fmt.Println("fact(5):", fact(5))
	fmt.Println()

	var fib func(n int) int

	fib = func(n int) int {
		if n < 0 {
			return 0
		} else if n < 2 {
			return n
		}

		return fib(n-1) + fib(n-2)
	}

	fmt.Println("fib(1):", fib(1))
	fmt.Println("fib(2):", fib(2))
	fmt.Println("fib(3):", fib(3))
	fmt.Println("fib(4):", fib(4))
	fmt.Println("fib(5):", fib(5))
	fmt.Println("fib(6):", fib(6))
}
