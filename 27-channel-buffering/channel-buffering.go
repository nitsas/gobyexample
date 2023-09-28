package main

import (
	"fmt"
	"time"
)

func write_and_log(c chan string, msg string) {
	fmt.Println("Writing \"", msg, "\" to channel", c)
	c <- msg
	fmt.Println("- Wrote \"", msg, "\"")
}

func main() {
	buffered_chan := make(chan string, 2)

	// fmt.Println(<-buffered_chan) // fatal error: all goroutines are asleep - deadlock!
	// Error if we try to read from an empty buffered channel

	write_and_log(buffered_chan, "1st buffered msg")
	write_and_log(buffered_chan, "2nd buffered msg")

	fmt.Println("Reading from buffered channel", buffered_chan)
	fmt.Println(<-buffered_chan)
	fmt.Println(<-buffered_chan)
	// fmt.Println(<-buffered_chan) // fatal error: all goroutines are asleep - deadlock!
	// Error if we try to read from an empty buffered channel
	fmt.Println()

	unbuffered_chan := make(chan string)

	// write_and_log(unbuffered_chan, "1st unbuffered msg") // fatal error: all goroutines are asleep - deadlock!
	// Error if we try to write to an unbuffered channel with no goroutine running

	go write_and_log(unbuffered_chan, "1st unbuffered msg")

	time.Sleep(time.Second)

	fmt.Println("Reading from unbuffered channel", unbuffered_chan)
	fmt.Println(<-unbuffered_chan)

	// time.Sleep(time.Second)
	// fmt.Println(<-unbuffered_chan) // fatal error: all goroutines are asleep - deadlock!
	// Error if we try to read from an unbuffered channel with no goroutine running
}
