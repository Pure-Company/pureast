package app

// UserService manages users
type UserService struct {
	users []User
}

// NewUserService creates a service
func NewUserService() *UserService {
	return &UserService{
		users: []User{},
	}
}

// AddUser adds a user
func (s *UserService) AddUser(u User) {
	if u.Validate() {
		s.users = append(s.users, u)
	}
}

// GetUser retrieves a user
func (s *UserService) GetUser(id int) *User {
	for _, u := range s.users {
		if u.ID == id {
			return &u
		}
	}
	return nil
}
