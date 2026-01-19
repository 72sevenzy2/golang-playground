package main

import "fmt"

type User struct {
	Name string
}

type Greeter interface {
	Print() string
}

func (u User) Print() string {
	return "user:" + u.Name
}

func show(p Greeter) {
	 fmt.Println(p.Print())
}

func main() {
	f := User{Name: "coolsaim"}
	show(f)
}
