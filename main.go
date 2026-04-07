package main

import "fmt"

func worker(tasks <-chan int, results chan<- int) {
	for i := range tasks {
		fmt.Println("recived with task number:", i)
		results <- i
	}
}

func main() {
	ch1 := make(chan int) // tasks
	ch2 := make(chan int) // results

	go func ()  { // workers
		for i := 0; i < 5; i++ {
			worker(ch1, ch2)
		}
		close(ch2)
	}()

	// senders
	go func ()  {
		for i := range 5 {
			ch1 <- i
		}
		close(ch1)
	}()

	for range 4 {
		<-ch2
	}
}