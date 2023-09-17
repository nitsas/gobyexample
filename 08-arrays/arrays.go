package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println("empty int array a:", a)
	a[4] = 100
	fmt.Println("set a[4] to 100, a:", a)
	fmt.Println("now a[4] is", a[4])
	fmt.Println("a has length", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("prefilled int array b:", b)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("two-dimensional array twoD:", twoD)
}
