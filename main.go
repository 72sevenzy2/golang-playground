package main

import "fmt"

func main() {
	var x int = 5
	var y int = 5

	result, err := add(x, y)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)

}

func add(a, b int) (int, error) {
	if a <= 0 || b <= 0 {
		return 0, fmt.Errorf("error")
	}
	return a + b, nil
}
