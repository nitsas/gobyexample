package main

import "fmt"

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("i:", i)

	zeroval(i)
	fmt.Println("zeroval(i)")
	fmt.Println("i:", i)

	zeroptr(&i)
	fmt.Println("zeroptr(i)")
	fmt.Println("i:", i)

	fmt.Println("&i:", &i)
}
