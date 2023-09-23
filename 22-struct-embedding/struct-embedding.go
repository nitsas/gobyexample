package main

import "fmt"

type identity struct {
	name string
}

func (id identity) identify() string {
	fmt.Println("- called identify() on:", id, "-")
	return "My name is" + id.name
}

type circle struct {
	identity
	radius float64
}

func main() {
	c := circle{
		radius:   5.0,
		identity: identity{name: "RedCircle"},
	}

	fmt.Println("circle c:", c)
	fmt.Println("c.radius:", c.radius)
	fmt.Println("c.name:", c.name)
	fmt.Println("c.identity.name:", c.identity.name)
	fmt.Println("c.identify():", c.identify())
	fmt.Println("c.identity.identify():", c.identity.identify())
	fmt.Println()

	type identifiable interface {
		identify()
	}

	idf := c
	fmt.Println("identifiable idf := c")
	fmt.Println("idf:", idf)
	fmt.Println("idf.identify():", idf.identify())
}
