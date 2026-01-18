package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4};

	b := a[1:3];

	fmt.Println("a:", a);
	fmt.Println("b:", b);

	b[0] = 99;

	fmt.Println("after change");
	fmt.Println("a:", a)
	fmt.Println("b:", b)

	c := make([]int, len(b))
	// fmt.Println(copy(c, b))
	copy(c, b)
	fmt.Println(c)
}