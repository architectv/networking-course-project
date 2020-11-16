package models

type User struct {
	Id string `json:"_id,omitempty" bson:"_id,omitempty"`
	Username string `json:"username" bson:"username"`
	FirstName string `json:"firstName,omitempty" bson:"firstName,omitempty"`
	LastName string `json:"lastName,omitempty" bson:"lastName,omitempty"`
	Email string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
	Phone string `json:"phone,omitempty" bson:"phone,omitempty"`
}
