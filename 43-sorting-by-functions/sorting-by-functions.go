package main

import (
	"cmp"
	"fmt"
	"slices"
)

func lenCmp(a string, b string) int {
	return cmp.Compare(len(a), len(b))
}

type Person struct {
	name string
	age  int
}

func makeListOfPeople() []Person {
	return []Person{
		{name: "Luffy", age: 19},
		{name: "Franky", age: 36},
		{name: "Robin", age: 31},
		{name: "Zoro", age: 21},
	}
}

func ageCmp(p1 Person, p2 Person) int {
	return cmp.Compare(p1.age, p2.age)
}

func ptrAgeCmp(p1 *Person, p2 *Person) int {
	return cmp.Compare(p1.age, p2.age)
}

func main() {
	fruits := []string{"watermelon", "banana", "apple", "orange", "kiwi"}
	fmt.Println("fruits:", fruits)
	fmt.Println("slices.SortFunc(fruits, lenCmp)")
	slices.SortFunc(fruits, lenCmp)
	fmt.Println("fruits:", fruits)
	fmt.Println()

	people := makeListOfPeople()

	fmt.Println("people:", people)
	fmt.Println("slices.SortFunc(people, ageCmp)")
	slices.SortFunc(people, ageCmp)
	fmt.Println("people:", people)
	fmt.Println()

	people = makeListOfPeople()
	peoplePtrs := make([]*Person, len(people))
	for i := 0; i < len(people); i++ {
		peoplePtrs[i] = &people[i]
	}

	fmt.Println("peoplePtrs:", peoplePtrs)
	fmt.Println("slices.SortFunc(peoplePtrs, ptrAgeCmp)")
	slices.SortFunc(peoplePtrs, ptrAgeCmp)
	fmt.Println("peoplePtrs:", peoplePtrs)
}
