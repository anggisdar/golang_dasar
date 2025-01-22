package main

import (
	"fmt"
	// "pertama/calculation"
	"pertama/management"
)

func main() {

	// fmt.Println("Halo, belajar Golang")
	// sentence := TestAja()

	// fmt.Println(sentence)

	// fmt.Println("Halo, belajar Golang")

	// result := calculation.Add(8, 9)

	// fmt.Println(result)
	// cara 1
	user := management.User{}
	user.ID = 1
	user.FirstName = "james"
	user.LastName = "bond"
	user.Email = "jamesbond@james.com"
	user.IsActive = true

	fmt.Println(user)

	//cara 2 horizontal mirip map
	userTwo := management.User{ID: 2, FirstName: "jack", LastName: "chan", Email: "jackchan@jack.co.id", IsActive: true}
	fmt.Println(userTwo)

	// cara 3 vertikal "urutannya bisa acak"
	userThree := management.User{
		Email:     "jacksparrow@jack.xyz",
		ID:        3,
		FirstName: "jack",
		LastName:  "sparrow",
		IsActive:  true,
	}

	fmt.Println(userThree)

	// cara 4 syarat harus id yang pertama
	user4 := management.User{4, "jack", "ma", "jackma@jack.org", true}
	user5 := management.User{2, "jack", "aa", "jackma@jack.org", true}
	// fmt.Println(userFour)

	// displayUser := displayUser(userFour)
	// fmt.Println(displayUser)

	//embedded struct
	users := []management.User{user4, user5}
	group := management.Group{"Gamer", user, users, true}

	displayGroup(group)
}

// embedded struct
func displayGroup(group management.Group) {
	fmt.Printf("Name: %s", group.Name)
	fmt.Println("")
	fmt.Printf("Member count : %d", len(group.Users))

}

// struct sebagai parameter
func displayUser(user management.User) string {
	return fmt.Sprintf("Name : %s %s, Email : %s", user.FirstName, user.LastName, user.Email)

}
