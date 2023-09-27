package main

import "fmt"

func main() {
	messages_chan := make(chan string)

	go func() { messages_chan <- "hello" }()

	msg := <-messages_chan

	fmt.Println("Received from channel messages_chan:", msg)
}
