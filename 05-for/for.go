package main

import "fmt"

func main() {
	fmt.Println("for i <= 3")
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	fmt.Println("\nfor j := 7; j <= 9; j++")
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	fmt.Println("\nfor { ... } with break")
	for {
		fmt.Println("loop and break")
		break
	}

	fmt.Println("\nfor n := 0; n <= 5; n++ { ... } continue if even num")
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
