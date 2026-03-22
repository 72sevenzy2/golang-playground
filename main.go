package main

import (
	"fmt"
	"sync"
)

func worker(id int, results chan<- int, jobs <-chan int, su *sync.WaitGroup) {
	defer su.Done()

	for i := range jobs{
		results <- i * 2
	}

	fmt.Println("workers done", id)
}

func main() {
	jobs := make(chan int)
	results := make(chan int)

	var b sync.WaitGroup

	for z := range 3 {
		b.Add(1)
		go worker(z, results, jobs, &b)
	}

	go func() {
		for i := range 3 {
			jobs <- i
		}
		close(jobs)
	}()

	go func ()  {
		b.Wait()
		close(results)
	}()

	for res := range results {
		fmt.Println("workers done: ", res)
	}
}