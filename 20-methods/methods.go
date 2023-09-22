package main

import "fmt"

type rect struct {
	width, height int
}

func (r rect) area() int {
	return r.width * r.height
}

func (r rect) doubleWidth() rect {
	r.width *= 2
	return r
}

func (rp *rect) increaseHeight() *rect {
	rp.height *= 2
	return rp
}

func main() {
	r := rect{width: 10, height: 20}
	rp := &r
	fmt.Println("r:", r)
	fmt.Println("rp:", rp)
	fmt.Println()

	fmt.Println("func area: receiver type rect, returns int, does NOT mutate")
	fmt.Println("r.area():", r.area())
	fmt.Println("rp.area():", rp.area())
	fmt.Println()

	fmt.Println("func doubleWidth: receiver type rect, returns new rect, cannot mutate")
	fmt.Println("r.doubleWidth():", r.doubleWidth())
	fmt.Println("r:", r)
	fmt.Println("rp.doubleWidth():", rp.doubleWidth())
	fmt.Println("r:", r)
	fmt.Println()

	fmt.Println("func increaseHeight: receiver type *rect, returns same rect, mutates")
	fmt.Println("r.increaseHeight:", r.increaseHeight())
	fmt.Println("r:", r)
	fmt.Println("rp.increaseHeight:", rp.increaseHeight())
	fmt.Println("r:", r)
	fmt.Println()
}
