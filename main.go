package main

import "fmt"

func main() {
	ch1 := make(chan int, 2)
	tasks := make(chan int)

	// sender
	go func() {
		for i := 0; i < 2; i++ {
			tasks <- i
		}
		close(tasks)
	}()

	// worker
	go func() {
		for i := range tasks {
			fmt.Println("received", i)
			ch1 <- i
		}
		close(ch1)
	}()

	// consumer
	for v := range ch1 {
		fmt.Println("result:", v)
	}
}