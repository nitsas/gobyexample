package main

import "fmt"

// Send the given message down the channel
func ping(pings chan<- string, msg string) {
	fmt.Println(msg)
	pings <- msg
}

// Pong a ping appropriately, and pass other messages through
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	if msg == "ping" {
		pongs <- "pong"
	} else if msg == "Ping!" {
		pongs <- "Pong!"
	} else if msg == "PING!" {
		pongs <- "PONG!"
	} else {
		pongs <- msg
	}
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	ping(pings, "ping")
	pong(pings, pongs)
	fmt.Println(<-pongs)
	fmt.Println()

	ping(pings, "Ping!")
	pong(pings, pongs)
	fmt.Println(<-pongs)
	fmt.Println()

	ping(pings, "PING!")
	pong(pings, pongs)
	fmt.Println(<-pongs)
	fmt.Println()

	ping(pings, "Hello world :)")
	pong(pings, pongs)
	fmt.Println(<-pongs)
	fmt.Println()
}
