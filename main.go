package main

import (
	"fmt"
	"time"
)

func worker(id int, ch1 <-chan int, ch2 chan<- int) {
	fmt.Println("go worker")
	for i := range ch1 {
		fmt.Println("received with id", id)
		ch2 <- i
	}
}

func main() {
	ch1 := make(chan int) // sender
	ch2 := make(chan int) // receiver

	for i := range 4 {
		time.Sleep(time.Second)
		go worker(i, ch1, ch2)
	}

	go func ()  {
		ch1 <- 1
		ch1 <- 2
		ch2 <- 3
	}()

	for range 3 {
		fmt.Println(<-ch2)
	}
}