package main

import (
	"fmt"
	"math"
)

type rect struct {
	width, height float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

type geometry interface {
	area() float64
	perim() float64
}

func measure(g geometry) {
	fmt.Println("measuring shape:", g)
	fmt.Println("  area:", g.area())
	fmt.Println("  perim:", g.perim())
}

func main() {
	c := circle{radius: 5.0}
	measure(c)

	r := rect{width: 4.0, height: 5.0}
	measure(r)
}
