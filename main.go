package main

import "fmt"

func main() {
	user_age := User{Age: 57}
	change(&user_age)
	fmt.Println(user_age.Age)
}

type User struct {
	Age int
}

func change(u *User) {
	u.Age = 53
}

