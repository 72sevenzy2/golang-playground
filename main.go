package main

import "fmt"

func send(ch chan<- int) {
	ch <- 6
}

func receive(ch <-chan int) {
	fmt.Println(<-ch)
}

func main() {
	ch := make(chan int)
	go send(ch)
	receive(ch)
}
