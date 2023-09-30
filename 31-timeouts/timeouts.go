package main

import (
	"fmt"
	"time"
)

func sleep_and_write(t time.Duration, c chan<- string, msg string) {
	time.Sleep(t)
	c <- msg
}

func main() {
	// Make a buffered channel so that the goroutine that writes to it won't
	// have to block, and we won't leak goroutines if the channel is never read.
	c1 := make(chan string, 1)
	go sleep_and_write(2*time.Second, c1, "- Result one after 2 secs!")

	fmt.Println("select with a timeout of 1 sec:")
	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(time.Second):
		fmt.Println("- Timeout after 1 sec")
	}
	fmt.Println()

	c2 := make(chan string, 1)
	go sleep_and_write(2*time.Second, c2, "- Result two after 2 secs!")

	fmt.Println("select with a timeout of 3 secs:")
	select {
	case res2 := <-c2:
		fmt.Println(res2)
	case <-time.After(3 * time.Second):
		fmt.Println("- Timeout after 3 secs")
	}
}
