package main

import "fmt"

func work(worker_id int, jobs <-chan int, done chan<- bool) {
	for {
		job, more := <-jobs
		if more {
			result := job * job * job
			fmt.Printf("worker %d: got job %d, result: %d\n", worker_id, job, result)
		} else {
			fmt.Printf("worker %d: no more jobs - terminating...\n", worker_id)
			done <- true
			return
		}
	}
}

func main() {
	// use a buffered channel for the jobs so that main can send
	// jobs without having to wait for a receiver
	jobs := make(chan int, 5)
	done1 := make(chan bool)
	done2 := make(chan bool)

	// spin up two workers just for fun
	go work(15, jobs, done1)
	go work(16, jobs, done2)

	for i := 0; i < 5; i++ {
		fmt.Printf("main: sending job %d\n", i)
		jobs <- i
	}

	// close the channel to indicate that we won't send anything more
	close(jobs)

	fmt.Println("main: waiting for the workers to complete")
	// Wait for the workers to complete the jobs before exiting
	<-done1
	<-done2
}
