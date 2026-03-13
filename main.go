package main

import "fmt"

func main() {
	var b [3]int

	for range len(b) {
		fmt.Println("hi")
	}
}