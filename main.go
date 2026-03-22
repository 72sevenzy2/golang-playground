package main

import (
	"fmt"
	"sync"
)

func worker(id int, results chan<- int, jobs <-chan int, su *sync.WaitGroup) {
	// this decreases the count by 1, so the loop i added is range 3 so the number of workers started will be 3 along with the count in the sync wg, but this will decrease it by 1 untill its 0 each time the func is called, which is 3 due to the loop.
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


	// gorountine to add 1 to the sync waitgroup and start the number of workers as needed
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
  

	// different gorountine to close the channel results ONLY after workers are done
	go func ()  {
		b.Wait()
		close(results)
	}()

	for res := range results {
		fmt.Println("workers done: ", res)
	}
}