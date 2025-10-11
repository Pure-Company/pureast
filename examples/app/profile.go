package app

// Address represents an address
type Address struct {
	Street string
	City   string
	Zip    string
}

// Profile combines User and Address
type Profile struct {
	User    User
	Address Address
	Active  bool
}

// NewProfile creates a new profile
func NewProfile(user User, addr Address) Profile {
	return Profile{
		User:    user,
		Address: addr,
		Active:  true,
	}
}

// IsComplete checks if profile is complete
func (p Profile) IsComplete() bool {
	return p.User.Validate() && p.Address.City != ""
}
