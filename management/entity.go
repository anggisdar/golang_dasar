package management

type User struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	IsActive  bool
}

// embedded struct
type Group struct {
	Name        string
	Admin       User
	Users       []User
	IsAvailable bool
}

// struct pointer
type Employee struct {
	ID         int
	Name, Role string
}
