package main

import "fmt"

type User struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	IsActive  bool
}

func main() {

	// cara 1
	user := User{}
	user.ID = 1
	user.FirstName = "james"
	user.LastName = "bond"
	user.Email = "jamesbond@james.com"
	user.IsActive = true

	fmt.Println(user)

	//cara 2 horizontal mirip map
	userTwo := User{ID: 2, FirstName: "jack", LastName: "chan", Email: "jackchan@jack.co.id", IsActive: true}
	fmt.Println(userTwo)

	// cara 3 vertikal "urutannya bisa acak"
	userThree := User{
		Email:     "jacksparrow@jack.xyz",
		ID:        3,
		FirstName: "jack",
		LastName:  "sparrow",
		IsActive:  true,
	}

	fmt.Println(userThree)

	// cara 4 syarat harus id yang pertama
	userFour := User{4, "jack", "ma", "jackma@jack.org", true}

	fmt.Println(userFour)
}
