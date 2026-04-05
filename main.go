package main

import (
	"fmt"
)

func main() {
	ch1 := make(chan int)

	go func ()  {
		ch1 <- 1
		ch1 <- 2
		ch1 <- 3
	}()

	for range 4 {
		fmt.Println(<-ch1)
	}
}