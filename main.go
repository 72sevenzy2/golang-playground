package main

import "fmt"

func divide(a, b int) (int, error) {
	if b == 0 || a == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}


func main() {
	// first exercise
	// a := []int{10, 20, 30, 40}
	// b := a[1:3]
	// b[0] = 99

	// c := make([]int, len(b))
	// copy(c, b)
	// b[1] = 100

	// fmt.Println(a) // 10, 99, 100, 40
	// fmt.Println(b) // 99, 100
	// fmt.Println(c) // 99, 30

	// second exercise
	result, err := divide(10, 5)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
}
