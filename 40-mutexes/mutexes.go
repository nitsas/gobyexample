package main

import (
	"fmt"
	"sync"
)

type Container struct {
	mu       sync.Mutex
	counters map[string]int
}

func (c *Container) inc(counterName string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counters[counterName]++
}

func incrementCounterTimes(c *Container, counterName string, times int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < times; i++ {
		c.inc(counterName)
	}
}

func main() {
	c := Container{
		counters: map[string]int{"a": 0, "b": 0},
	}
	fmt.Println("main: c.counters:", c.counters)

	var wg sync.WaitGroup

	fmt.Println("main: Starting 3 goroutines that will increment the counters")
	wg.Add(3)
	go incrementCounterTimes(&c, "a", 10000, &wg)
	go incrementCounterTimes(&c, "a", 10000, &wg)
	go incrementCounterTimes(&c, "b", 10000, &wg)

	fmt.Println("main: Waiting for goroutines to finish")
	wg.Wait()

	fmt.Println("main: c.counters:", c.counters)
}
