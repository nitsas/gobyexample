package main

import (
	"fmt"
	"time"
)

type JobResult struct {
	job    int
	result int
}

func work(id int, jobs <-chan int, jobResults chan<- JobResult) {
	for job := range jobs {
		fmt.Printf("worker %d: got job %d, working ...\n", id, job)
		time.Sleep(time.Second)
		result := job * job * job
		fmt.Printf("worker %d: job %d done!\n", id, job)
		jobResults <- JobResult{job: job, result: result}
	}
}

func main() {
	const numJobs = 5
	const numWorkers = 3

	jobs := make(chan int, numJobs)
	jobResults := make(chan JobResult, numJobs)

	fmt.Printf("main: starting %d workers\n", numWorkers)
	for i := 1; i <= numWorkers; i++ {
		go work(i, jobs, jobResults)
	}

	fmt.Printf("main: enqueueing %d jobs\n", numJobs)
	for i := 1; i <= numJobs; i++ {
		jobs <- i
	}
	fmt.Println("main: closing channel jobs")
	close(jobs)

	for i := 1; i <= numJobs; i++ {
		jobResult := <-jobResults
		fmt.Printf("main: job %d: result: %d\n", jobResult.job, jobResult.result)
	}
}
