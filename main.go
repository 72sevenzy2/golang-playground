package main

import "fmt"

func worker(results chan<- int, jobs <-chan int) {
	for i := range jobs {
		results <- i
	}
}

func main() {
	jobs := make(chan int)
	results := make(chan int)

	go worker(results, jobs)

	go func ()  {
		for i := range 2 {
			jobs <- i
		}
		close(jobs)
	}()

	fmt.Println(<-results)
	fmt.Println(<-results)
}