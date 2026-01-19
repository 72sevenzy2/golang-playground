package main

import "fmt"

func main() {
	// first exercise
	a := []int{10, 20, 30, 40}
	b := a[1:3]
	b[0] = 99

	c := make([]int, len(b))
	copy(c, b)
	b[1] = 100

	fmt.Println(a) // 10, 99, 100, 40
	fmt.Println(b) // 99, 100
	fmt.Println(c) // 99, 30
}
