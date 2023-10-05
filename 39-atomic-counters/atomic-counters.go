package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func incrementCount(countAtomic *atomic.Uint64, countInt *int, times int, wg *sync.WaitGroup) {
	for i := 0; i < times; i++ {
		countAtomic.Add(1)
		*countInt += 1
	}

	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	var countAtomic atomic.Uint64
	countInt := 0

	fmt.Println("main: starting workers")
	numWorkers := 50
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go incrementCount(&countAtomic, &countInt, 1000, &wg)
	}

	fmt.Println("main: waiting for workers to finish")
	wg.Wait()

	fmt.Println("Done! Results:")
	fmt.Println("- countAtomic:", countAtomic.Load())
	fmt.Println("- countInt:", countInt)
	// We expect countInt to be lower than countAtomic, because the goroutines
	// will run into race conditions and some increments will be lost.
}
