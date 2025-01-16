package main

import "fmt"

type employee struct {
	name string
	age  int
}

type User struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	IsActive  bool
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

	user := User{}
	user.ID = 1
	user.FirstName = "james"
	user.LastName = "bond"
	user.Email = "jamesbond@james.com"
	user.IsActive = true

	fmt.Println(user)

}
