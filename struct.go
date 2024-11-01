package main

import "fmt"

type employee struct {
	name string
	age  int
}

func main() {
	employee1 := employee{
		name: "ujek",
		age:  24,
	}

	employee2 := employee{}
	employee2.name = "arga"
	employee2.age = 23

	fmt.Println("Name: ", employee1.name)
	fmt.Println("age: ", employee1.age)
	fmt.Println("--------------------------")
	fmt.Println("Name: ", employee2.name)
	fmt.Println("age: ", employee2.age)
}
