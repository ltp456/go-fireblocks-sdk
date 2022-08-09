package types

type AllUsers struct {
	Users []User `json:"users"`
}

type User struct {
	ID        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Role      string `json:"role"`
	Email     string `json:"email"`
	Enabled   bool   `json:"enabled"`
}
