package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)

	go func ()  {
		ch1 <- 1
		ch1 <- 2
		ch1 <- 3
	}()

	for range 3 {
		time.Sleep(time.Second)
		fmt.Println(<-ch1)
	}
}