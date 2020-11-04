package models

type User struct {
	Id string `json:"_id,omitempty"`
	Username string `json:"username"`
	FirstName string `json:"firstName,omitempty"`
	LastName string `json:"lastName,omitempty"`
	Email string `json:"email"`
	Password string `json:"password"`
	Phone string `json:"phone,omitempty"`
}
