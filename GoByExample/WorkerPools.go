package main

import (
	"fmt"
	"time"
)

func worker(w int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("Worker", w, "started job", j)
		time.Sleep(1 * time.Second)
		fmt.Println("Worker", w, "finished job", j)
		results <- j * 2
	}
}

func main() {

	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for i := 0; i < 3; i++ {
		go worker(i+1, jobs, results)
	}

	for i := 0; i < 5; i++ {
		jobs <- i + 1
	}
	close(jobs)

	for i := 1; i <= 5; i++ {
		fmt.Println(<-results)
	}
}
