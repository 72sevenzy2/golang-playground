package main

import (
	"fmt"
	"sync"
)

func worker(wg *sync.WaitGroup, tasks <-chan int, results chan<- int) {
	defer wg.Done()
	for i := range tasks {
		fmt.Println("received task number:", i)
		results <- i * 2 // example tasks
	}
}

func main() {
	tasks := make(chan int, 5)
	results := make(chan int, 5) // buffered channels = queue

	var wg sync.WaitGroup

	for range 3 {
		wg.Add(1)
		go worker(&wg, tasks, results)
	}

	go func() {
		for i := range 3 {
			tasks <- i
		}
		close(tasks)
	}()

	go func ()  {
		wg.Wait()
		close(results)
	}()

	for range 3 {
		fmt.Println(<-results)
	}
}
