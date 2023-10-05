package main

import (
	"fmt"
	"time"
)

func sampleRequests(howMany int) chan int {
	requests := make(chan int, howMany)
	for i := 0; i < howMany; i++ {
		requests <- i + 1
	}

	return requests
}

func work(requests <-chan int, done chan<- bool) {
	for req := range requests {
		response := req * req * req
		fmt.Println("worker: Got request", req, "at", time.Now().Format(time.RFC3339Nano), "- response:", response)
	}
	done <- true
}

func rateLimit(requests <-chan int, rate time.Duration) chan int {
	rateLimitedRequests := make(chan int)

	ticker := time.NewTicker(rate)
	go func() {
		for req := range requests {
			<-ticker.C
			rateLimitedRequests <- req
		}
		close(rateLimitedRequests)
	}()

	return rateLimitedRequests
}

func rateLimitBursty(requests <-chan int, rate time.Duration, burstCapacity int) chan int {
	rateLimitedRequests := make(chan int, burstCapacity)
	burstyTicker := make(chan time.Time, burstCapacity)

	for i := 0; i < burstCapacity; i++ {
		burstyTicker <- time.Now()
	}
	go func() {
		for t := range time.Tick(rate) {
			burstyTicker <- t
		}
	}()

	go func() {
		for req := range requests {
			<-burstyTicker
			rateLimitedRequests <- req
		}
		close(rateLimitedRequests)
	}()

	return rateLimitedRequests
}

func main() {
	numRequests := 5
	rate := 400 * time.Millisecond

	requests := sampleRequests(numRequests)
	close(requests)

	fmt.Println("Limiter allowing requests every", rate)
	rateLimitedRequests := rateLimit(requests, rate)
	done := make(chan bool)
	go work(rateLimitedRequests, done)
	<-done
	fmt.Println()

	burstCapacity := 3
	fmt.Println("Limiter allowing bursts of", burstCapacity, "requests, every", rate)
	requests2 := sampleRequests(numRequests)
	requests2Bursty := rateLimitBursty(requests2, rate, burstCapacity)
	go work(requests2Bursty, done)

	// Now wait a bit and then send another burst of requests
	// NOTE: When we sleep for 2 seconds this works as expected: It allows
	//       a second burst of 3 requests and then more requests every 400ms.
	//       But if we sleep for 3+ seconds, then all requests are executed
	//       in quick succession. It's as if the bursty limiter gets filled-in
	//       beyond its capacity. :/
	//       The same happens with the original example code.
	// TODO: Revisit this later.
	time.Sleep(2 * time.Second)
	requests2 <- 10
	requests2 <- 20
	requests2 <- 30
	requests2 <- 40
	requests2 <- 50
	close(requests2)

	// Wait for the worker to finish
	<-done
}
