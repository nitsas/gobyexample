package main

import "fmt"

func main() {
	queue := make(chan string, 2)

	fmt.Println("Writing to channel queue")
	queue <- "first msg"
	queue <- "second msg"
	fmt.Println("Closing channel queue")
	close(queue)
	fmt.Println()

	fmt.Println("Iterating over channel queue:")
	for msg := range queue {
		fmt.Println(msg)
	}
}
