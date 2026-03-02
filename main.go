	package main

	import "fmt"

	func main() {
		var a int = 1;
		var b int = 12;

		result, err := divide(a, b);

		if err != nil {
			fmt.Println(err.Error())
		}

		fmt.Println(result)

	}