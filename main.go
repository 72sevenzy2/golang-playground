package main

import "fmt"

var arr [3][2]int

func main() {
	arr[1][0] = 99
	fmt.Println(arr[1][0])
	fmt.Println(arr)

	// seperate from this exercise ^^

	var i counter = 0

	var b incrementer = &i

	b.increase()
	
}