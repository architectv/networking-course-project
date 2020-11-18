package models

type User struct {
	Id       string `json:"_id,omitempty" bson:"_id,omitempty"`
	Nickname string `json:"nickname" bson:"nickname" valid:"length(3|32)"`
	Email    string `json:"email" bson:"email" valid:"email"`
	Avatar   string `json:"avatar" bson:"avatar"`
	Password string `json:"password" bson:"password" valid:"length(6|32)"`
}
