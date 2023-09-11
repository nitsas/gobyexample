package main

import "fmt"

func main() {
	var a = "initial"
	fmt.Println("type inference:")
	fmt.Println("var a =", a)

	var b, c int = 1, 2
	fmt.Println("initializing multiple variables:")
	fmt.Println("var b, c int =", b, ",", c)

	var d = true
	fmt.Println("boolean variable: ", d)

	var e int
	fmt.Println("zero initializing an int: ", e)

	f := "apple"
	fmt.Println("declaration and initialization with := inside functions:")
	fmt.Println(f)
}
