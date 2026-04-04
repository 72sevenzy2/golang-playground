package main

import "fmt"

	func worker(id int, tasks <-chan int, results chan<- int) {
		for i := range tasks {
			fmt.Println("id", id)
			results <- i
		}
	}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	for i := range 3 {
		go worker(i, ch2, ch1)
	}

	go func ()  {
		ch2 <- 1
		ch2 <- 2
		ch2 <- 3
	}()

	for range 3 {
		fmt.Println(<-ch1)
	}
}