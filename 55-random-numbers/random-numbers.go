package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("rand.Intn(100):", rand.Intn(100), rand.Intn(100))
	fmt.Println("rand.Int():", rand.Int(), rand.Int())
	fmt.Println()

	fmt.Println("rand.Float64():", rand.Float64(), rand.Float64())
	fmt.Println("100*rand.Float64():", 100*rand.Float64(), 100*rand.Float64())
	fmt.Println("rand.NormFloat64():", rand.NormFloat64(), rand.NormFloat64(), rand.NormFloat64())
	fmt.Println()

	fmt.Println("rand.Perm(10):", rand.Perm(10))
	fmt.Println()

	fruits := []string{"apple", "banana", "peach", "strawberry", "watermelon"}
	fmt.Println("fruits:", fruits)
	swapFruits := func(i int, j int) {
		temp := fruits[j]
		fruits[j] = fruits[i]
		fruits[i] = temp
	}
	fmt.Println("rand.Shuffle(len(fruits), swapFruits)")
	rand.Shuffle(len(fruits), swapFruits)
	fmt.Println("fruits:", fruits)
	fmt.Println()

	// source1 := rand.NewSource(time.Now().UnixNano())
	source1 := rand.NewSource(int64(time.Now().Day()))
	fmt.Println("source1 := rand.NewSource(<int64>)")
	r := rand.New(source1)
	fmt.Println("r := rand.New(source1)")
	fmt.Println("r.Intn(100):", r.Intn(100), r.Intn(100), r.Intn(100), r.Intn(100), r.Intn(100))

	// RNG with the same seed:
	source2 := rand.NewSource(int64(time.Now().Day()))
	fmt.Println("source2 := rand.NewSource(<same-seed>)")
	r2 := rand.New(source2)
	fmt.Println("r2 := rand.New(source2)")
	fmt.Println("r2.Intn(100):", r2.Intn(100), r2.Intn(100), r2.Intn(100), r2.Intn(100), r2.Intn(100))
}
