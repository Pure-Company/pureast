package app

import "fmt"

// User represents a user
type User struct {
	ID    int
	Name  string
	Email string
}

// NewUser creates a new user
func NewUser(id int, name string) User {
	return User{
		ID:   id,
		Name: name,
	}
}

// Print prints the user
func (u User) Print() {
	fmt.Printf("User: %s (ID: %d)\n", u.Name, u.ID)
}

// Validate validates the user
func (u User) Validate() bool {
	return u.ID > 0 && u.Name != ""
}
