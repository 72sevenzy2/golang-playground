package main

import (
	"fmt"
)

// worker function
func worker(id int, tasks <-chan int, results chan<- int) {
	for task := range tasks {
		fmt.Println("Worker", id, "processing", task)
		results <- task * 2 // example processing
	}
}

func main() {
	tasks := make(chan int, 5)
	results := make(chan int, 5)

	// start 3 workers
	for i := 1; i <= 3; i++ {
		go worker(i, tasks, results)
	}

	tasks <- 1
	tasks <- 2
	close(tasks)

	// collect results
	for i := 1; i <= 2; i++ {
		fmt.Println("Result:", <-results)
	}
}