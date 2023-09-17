package main

import (
	"fmt"
	"maps"
)

func main() {
	m := make(map[string]int)
	fmt.Println("empty map m:", m)

	m["k1"] = 10
	m["k2"] = 20
	fmt.Println("set some parts of map m:", m)
	fmt.Println("get nonexisting key m[\"k3\"]", m["k3"])
	fmt.Println("len(m):", len(m))
	v2 := m["k2"]
	fmt.Println("v2 := m[\"k2\"]:", v2)
	fmt.Println()

	delete(m, "k2")
	fmt.Println("deleted key \"k2\" from m:", m)
	fmt.Println("variable v2:", v2)
	clear(m)
	fmt.Println("cleared m:", m)
	_, present := m["k2"]
	fmt.Println("check if key \"k2\" is present in m:", present)
	fmt.Println()

	m2 := map[string]int{"k4": 40, "k5": 50}
	fmt.Println("prefilled map m2:", m2)
	m3 := map[string]int{"k4": 40, "k5": 50}
	fmt.Println("prefilled map m3:", m3)
	fmt.Println("maps.Equal(m3, m4):", maps.Equal(m2, m3))
}
