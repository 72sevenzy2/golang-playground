package main

import (
	"fmt"
	"sync"
)

type task struct {
	index int
	value int 
}

type result struct {
	index int
	value int
}

func worker(wg *sync.WaitGroup, tasks <-chan task, results chan<- result) {
	defer wg.Done()
	for i := range tasks {
		fmt.Println("received task number:", i)
		results <- result{index: i.index, value: i.value * 2} // example tasks
	}
}

func main() {
	tasks := make(chan task, 5)
	results := make(chan result, 5) // buffered channels = queue

	var wg sync.WaitGroup

	for range 3 {
		wg.Add(1)
		go worker(&wg, tasks, results)
	}

	go func() {
		for i := range 3 {
			tasks <- task{index: i, value: i}
		}
		close(tasks)
	}()

	go func ()  {
		wg.Wait()
		close(results)
	}()

	output := make([]int, 3)

	for v := range results {
		output[v.index] = v.value
	}

	for _, v := range output {
		fmt.Println(v)
	}
}
