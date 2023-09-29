package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Print("Working... ")
	time.Sleep(time.Second)
	fmt.Println("Done!")

	done <- true
}

func main() {
	// It works with an unbuffered channel too.
	done := make(chan bool, 1)
	go worker(done)

	<-done
}
