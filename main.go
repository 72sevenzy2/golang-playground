package main

import (
	"fmt"
)

type incrementer interface {
	increment()
}

type counter int

func (c *counter) increment() {
	*c++
}

func main() {
	var i incrementer

	c := counter(0)
	i = &c
	i.increment()
	fmt.Println(c)
}
