package main

import (
	"fmt"
	"math/rand"
	"time"
)

type readOp struct {
	key  int
	resp chan int
}

type writeOp struct {
	key  int
	val  int
	resp chan bool
}

func runStore(
	killSwitch <-chan bool,
	reads <-chan readOp, countReads *int,
	writes <-chan writeOp, countWrites *int,
) {
	state := make(map[int]int)
	for {
		select {
		case <-killSwitch:
			fmt.Println("store: Stopped")
			return
		case read := <-reads:
			read.resp <- state[read.key]
			*countReads++
		case write := <-writes:
			state[write.key] = state[write.val]
			write.resp <- true
			*countWrites++
		}
	}
}

func runReader(reads chan<- readOp) {
	for {
		key := rand.Intn(5)
		resp := make(chan int)
		reads <- readOp{key: key, resp: resp}
		<-resp
		time.Sleep(time.Millisecond)
	}
}

func runWriter(writes chan<- writeOp) {
	for {
		key := rand.Intn(5)
		val := rand.Intn(100)
		resp := make(chan bool)
		writes <- writeOp{key: key, val: val, resp: resp}
		<-resp
		time.Sleep(time.Millisecond)
	}
}

func main() {
	numReaders := 100
	numWriters := 10

	storeKillSwitch := make(chan bool)
	reads := make(chan readOp)
	writes := make(chan writeOp)

	fmt.Println("main: Starting the store")
	countReads := 0
	countWrites := 0
	go runStore(storeKillSwitch, reads, &countReads, writes, &countWrites)

	for i := 0; i < numReaders; i++ {
		go runReader(reads)
	}

	for i := 0; i < numWriters; i++ {
		go runWriter(writes)
	}

	fmt.Println("main: Sleeping for a couple seconds...")
	time.Sleep(2 * time.Second)

	fmt.Println("main: Killing the store...")
	storeKillSwitch <- true
	fmt.Println()

	fmt.Println("main: number of reads in the store:", countReads)
	fmt.Println("main: number of writes in the store:", countWrites)
}
