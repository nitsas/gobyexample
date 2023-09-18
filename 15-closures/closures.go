package main

import "fmt"

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	fmt.Println("nextInt() := intSeq()")
	nextInt := intSeq()
	fmt.Println("nextInt():", nextInt())
	fmt.Println("nextInt():", nextInt())
	fmt.Println("nextInt():", nextInt())

	fmt.Println("nextInt2() := intSeq()")
	nextInt2 := intSeq()
	fmt.Println("nextInt2():", nextInt2())
	fmt.Println("nextInt2():", nextInt2())
}
