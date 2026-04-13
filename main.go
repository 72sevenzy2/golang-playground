package main

import (
	"fmt"
)



// fan in pattern

func worker(id int, res chan<- string) {
	for range 3 {
		res <- fmt.Sprintf("received worker with id %d", id)
	}
	close(res)
}

func main() {
	ch1 := make(chan string) // receiver
	ch2 := make(chan string) // tasks

	go worker(1, ch1)
	go worker(1, ch2)
}