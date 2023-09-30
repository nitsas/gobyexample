package main

import (
	"fmt"
	"time"
)

func print_when_timer(t *time.Timer, timer_name string, msg string) {
	<-t.C
	fmt.Println("Timer", timer_name, "fired:", msg)
}

func main() {
	fmt.Println("Creating timer t1 that will fire in 2 secs")
	t1 := time.NewTimer(2 * time.Second)

	// block until the timer fires
	print_when_timer(t1, "t1", "Hello!")
	fmt.Println()

	fmt.Println("Creating timer t2 that will fire in 1 sec")
	t2 := time.NewTimer(time.Second)
	// wait on the timer asynchronously
	go print_when_timer(t2, "t2", "World!")
	fmt.Println("Stopping timer t2")
	stop2 := t2.Stop()
	if stop2 {
		fmt.Println("Stopped timer t2!")
	}
	fmt.Println()

	fmt.Println("Wait for 2 secs until all goroutines finish")
	time.Sleep(2 * time.Second)
}
