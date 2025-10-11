package example

import "fmt"

// User represents a user
type User struct {
    ID   int
    Name string
    Email string
}

// Address represents an address
type Address struct {
    Street string
    City   string
}

// Profile combines User and Address
type Profile struct {
    User    User
    Address Address
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

// GetName returns the user's name
func (u User) GetName() string {
    return u.Name
}

// SetEmail sets the email
func (u *User) SetEmail(email string) {
    u.Email = email
}

