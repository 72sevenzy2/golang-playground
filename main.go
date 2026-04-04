package main

import "fmt"

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func ()  {
		ch1 <- 1
	}()

	go func ()  {
		ch2 <- 2
	}()

	go func ()  {
		ch1 <- 4
	}()

	go func ()  {
		ch1 <- 7
	}()

	go func ()  {
		ch2 <- 3
	}()

	for range 2 {
		fmt.Println(<-ch1)
		fmt.Println(<-ch2)
	}
}