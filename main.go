// simple worker pool

package main

import (
	"fmt"
	"time"
)

func worker(id int, tasks <-chan int, results chan<- int) {
	for i := range tasks {
		fmt.Println("received task from worker", id)
		results <- i
	}
}

func main() {
	tasksch := make(chan int)
	resch := make(chan int)

	for i := range 3 {
		time.Sleep(time.Second)
		go worker(i, tasksch, resch)
	}

	tasksch <- 1
	close(tasksch)

}
