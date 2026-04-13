package main

import (
	"fmt"
	"sync"
)

// fan in pattern

func worker(id int, res chan<- string) {
	for range 3 {
		res <- fmt.Sprintf("received worker with id %d", id)
	}
	close(res)
}

func fanin(wg *sync.WaitGroup, channels ...<-chan string) <-chan string {
	out := make(chan string)

	wg.Add(len(channels))
	for _, value := range channels {
		go func (ch <-chan string)  {
			defer wg.Done()
			for v := range ch {
				out <- v
			}
		}(value)
	}

	go func ()  {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {
	ch1 := make(chan string) // receiver
	ch2 := make(chan string) // tasks

	go worker(1, ch1)
	go worker(1, ch2)
}
