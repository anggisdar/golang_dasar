package main

import "fmt"

type User struct {
	Name, Role string
	IsActive   bool
}

func (user User) sayHello() {
	fmt.Println("Hello, My name is", user.Name, "my role", user.Role, true)
}

func main() {
	anggis := User{Name: "Anggis", Role: "Backend Devloper Enthusiast"}
	anggis.sayHello()
}
