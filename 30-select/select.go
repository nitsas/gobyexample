// Run this with `time go run select.go` to see that the execution time
// is only 2 secs, because both goroutines execute concurrently.
package main

import (
	"fmt"
	"time"
)

func sleep_and_go(c chan<- string, duration int, msg string) {
	fmt.Printf("sleep_and_go for %ds will write %s\n", duration, msg)
	time.Sleep(time.Duration(duration) * time.Second)
	c <- msg
}

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go sleep_and_go(c2, 2, "two")
	go sleep_and_go(c1, 1, "one")

	for i := 0; i < 2; i++ {
		select {
		case msg := <-c1:
			fmt.Printf("Got: %s\n", msg)
		case msg := <-c2:
			fmt.Printf("Got: %s\n", msg)
		}
	}
}
