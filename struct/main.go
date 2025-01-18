package main

import "fmt"

type User struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	IsActive  bool
}

//embedded struct
type Group struct {
	Name        string
	Admin       User
	Users       []User
	IsAvailable bool
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
	user4 := User{4, "jack", "ma", "jackma@jack.org", true}
	user5 := User{2, "jack", "aa", "jackma@jack.org", true}
	// fmt.Println(userFour)

	// displayUser := displayUser(userFour)
	// fmt.Println(displayUser)

	//embedded struct
	users := []User{user4, user5}
	group := Group{"Gamer", user, users, true}

	displayGroup(group)
}

//embedded struct
func displayGroup(group Group) {
	fmt.Printf("Name: %s", group.Name)
	fmt.Println("")
	fmt.Printf("Member count : %d", len(group.Users))

}

// struct sebagai parameter
func displayUser(user User) string {
	result := fmt.Sprintf("Name : %s %s, Email : %s", user.FirstName, user.LastName, user.Email)
	return result
}
