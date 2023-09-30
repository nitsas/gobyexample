package main

import "fmt"

func main() {
	// Use unbuffered channels which block on send/receive
	messages := make(chan string)
	signals := make(chan string)

	fmt.Println("select with a default case trying a non-blocking receive from messages:")
	select {
	case msg := <-messages:
		fmt.Println("- received message:", msg)
	default:
		fmt.Println("- no sender ready to send")
	}
	fmt.Println()

	fmt.Println("select with a default case trying a non-blocking send on messages:")
	outgoing_msg := "Hi!"
	select {
	case messages <- outgoing_msg:
		fmt.Println("- sent", outgoing_msg)
	default:
		fmt.Println("- no receiver ready to receive")
	}
	fmt.Println()
	// The message was not sent now and also won't be sent later
	// when we try to read from messages.

	fmt.Println("multi-clause select with default case trying non-blocking receives from messages and signals:")
	select {
	case msg2 := <-messages:
		fmt.Println("- got message", msg2)
	case sig := <-signals:
		fmt.Println("- got signal", sig)
	default:
		fmt.Println("- no sender ready on messages and no receiver ready on signals")
	}
}
