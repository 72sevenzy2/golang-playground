package main

import (
	"fmt"
	"time"
)

func worker(id int, rec chan<- int, sen <-chan int) {
	fmt.Println("worker id", id)
	for i := range sen {
		fmt.Println("received with id", i)
		rec <- i
	}
}

func main() {
	ch1 := make(chan int)
	rec := make(chan int)

	for i := 0; i < 3; i++ {
		time.Sleep(time.Second)
		go worker(i, rec, ch1)
	}

	go func ()  {
		ch1 <- 1
		ch1 <- 2
		ch1 <- 3
		close(ch1)
	}()

	for i := 1; i < 4; i++ {
		fmt.Println(<-rec)
	}

}