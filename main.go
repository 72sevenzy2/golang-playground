package main

import (
	"fmt"
	"sync"
)

func worker(wg *sync.WaitGroup, id int, results chan<- int, tasks <-chan int) {
	defer wg.Done()
	fmt.Println("received work with id", id)
	for i := range tasks {
		fmt.Println("task id", i)
		results <- i
	}
}

func main() {
	tasks := make(chan int)
	results := make(chan int)

	var wg sync.WaitGroup

	for i := range 5 {
		wg.Add(1)
		go worker(&wg, i, results, tasks)
	}

	go func ()  {
		for i := range 3 {
			tasks <- i
		}
		close(tasks)
	}()

	go func ()  {
		wg.Wait()
		close(results)
	}()

	for range results {
		fmt.Println(<-results)
	}
}