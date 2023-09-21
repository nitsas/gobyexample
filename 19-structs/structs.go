package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	np := person{name: name}
	np.age = 42
	return &np
}

func main() {
	fmt.Println("person{\"Bob\", 20}:", person{"Bob", 20})
	fmt.Println("person{age: 21, name: \"Luffy\"}:", person{age: 21, name: "Luffy"})
	fmt.Println("person{name:\"Fred\"}:", person{name: "Fred"})
	fmt.Println("person{age: 99}:", person{age: 99})
	fmt.Println()

	fmt.Println("&person{name: \"Ann\", age: 40}:", &person{name: "Ann", age: 40})
	fmt.Println()

	p := person{name: "Sean", age: 50}
	fmt.Println("p:", p)
	fmt.Println("p.name:", p.name)

	pp := &p
	fmt.Println("pp = &p:", pp)
	fmt.Println("pp.age:", pp.age)
	fmt.Println()

	p.name = "Seann"
	fmt.Println("p.name = \"Seann\"")
	pp.age = 51
	fmt.Println("pp.age = 51")
	fmt.Println("p:", p)
	fmt.Println()

	dog := struct {
		name   string
		isGood bool
	}{"Rex", true}
	fmt.Println("single-use struct, for variable dog:", dog)
}
