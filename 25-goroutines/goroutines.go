package main

import (
	"fmt"
	"time"
)

func printThreeTimes(msg string) {
	for i := 0; i < 3; i++ {
		fmt.Println(msg, ":", i)
	}
}

func main() {
	fmt.Println("Playing around with functions and goroutines:")

	printThreeTimes("direct function call")

	go printThreeTimes("named goroutine")

	go func(msg string) {
		for i := 0; i < 4; i++ {
			fmt.Println(msg, ":", i)
		}
	}("unnamed inline goroutine")

	time.Sleep(time.Second)
	fmt.Println("done!")
}
