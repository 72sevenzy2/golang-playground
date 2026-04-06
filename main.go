package main

import "fmt"

func main() {
	ch1 := make(chan int, 2) // buffered channels, size 2
	tasks := make(chan int)

	go func ()  {
		tasks <- 1
		tasks <- 2
	}()

	for i := range tasks {
		ch1 <- i
		fmt.Println("received", i)
	}
}