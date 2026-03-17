package main

import "fmt"

func main() {
	ch := make(chan int)

	ch <- 6

	b := <- ch

	fmt.Println(b)

}
