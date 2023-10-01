package main

import (
	"fmt"
	"sync"
	"time"
)

func work(id int) {
	fmt.Printf("worker %d: starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("worker %d: finishing\n", id)
}

type workFunc func(int)

func runInWaitgroup(work workFunc, wg *sync.WaitGroup, id int) {
	// Cannot move wg.Add(1) here because main may reach wg.Wait()
	// before this goroutine runs wg.Add(1).
	defer wg.Done()
	work(id)
}

func main() {
	var wg sync.WaitGroup
	const numWorkers = 5

	fmt.Println("main: running workers")
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go runInWaitgroup(work, &wg, i+1)
	}

	fmt.Println("main: waiting for workers to finish...")
	wg.Wait()
	fmt.Println("main: exiting")
}
