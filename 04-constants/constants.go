package main

import (
	"fmt"
	"math"
)

const s string = "constant"

func main() {
	fmt.Println("using a const:", s)

	const n = 500000000
	fmt.Println("Declared a numeric const without a type:", n)

	const d = 3e20 / n
	fmt.Println("const result of division:", d)

	fmt.Println("giving it a type of int64:", int64(d))

	fmt.Println("math.Sin() expects a float64, so it typecasts the untyped const:", math.Sin(n))
}
