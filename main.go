package main

import "fmt"

// func divide(a, b int) (int, error) {
// 	if b == 0 || a == 0 {
// 		return 0, fmt.Errorf("cannot divide by zero")
// 	}
// 	return a / b, nil
// }

// exercise 4
type Greeter interface {
	greet() string
}

type Person struct {
	Name string
}

func (p *Person) greet() string{
	 return "person name " + p.Name
}

type Bot struct {
	ID int
}

func (b Bot) greet() string{
	return fmt.Sprintf("bot id", b.ID)
}

func sayHello(g Greeter) {
	fmt.Println(g.greet())
}

func main() {
	p := Person{Name: "cool"}
	a := Bot{ID: 5}

	sayHello(&p)
	sayHello(a)

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
	// result, err := divide(10, 5)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(result)

	// var i Incrementer

	// c := Counter(5)
	// i = &c
	// i.int()
	// fmt.Println(c)
}

// exercise 3
// type Counter int

// type Incrementer interface {
// 	int()
// }

// func (c *Counter) int() {
// 	*c++
// }
