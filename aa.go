package main

import "fmt"

type counter int

type incrementer interface {
	increase()
}

func (s *counter) increase() {
	*s = *s + 1
	fmt.Println(*s)
}
