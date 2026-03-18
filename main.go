package main

import "fmt"

func worker(jobs <-chan int, results chan<- int) {
	for job := range jobs {
		results <- job * 2
	}
}

func main() {
	jobs := make(chan int)
	results := make(chan int)

	// start worker
	go worker(jobs, results)

	// send jobs
	jobs <- 1
	jobs <- 2

	// receive ONLY ONE result
	fmt.Println(<-results)
}