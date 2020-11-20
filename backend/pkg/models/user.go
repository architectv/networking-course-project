package models

type User struct {
	Id       int    `json:"id,omitempty"`
	Nickname string `json:"nickname" valid:"length(3|32)"`
	Email    string `json:"email" valid:"email"`
	Password string `json:"password" valid:"length(6|32)"`
	Avatar   string `json:"avatar"`
}
