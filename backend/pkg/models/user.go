package models

type User struct {
	Id        int    `json:"id,omitempty"`
	Nickname  string `json:"nickname" valid:"length(3|32)"`
	Firstname string `json:"firstname" valid:"length(1|32)"`
	Lastname  string `json:"lastname" valid:"length(1|32)"`
	Email     string `json:"email" valid:"email"`
	Phone     string `json:"phone" valid:"numeric"`
	Password  string `json:"password" valid:"length(6|32)"`
	Avatar    string `json:"avatar"`
}

type UpdateUser struct {
	Nickname  *string `json:"nickname" valid:"length(3|32)"`
	Firstname *string `json:"firstname" valid:"length(1|32)"`
	Lastname  *string `json:"lastname" valid:"length(1|32)"`
	Email     *string `json:"email" valid:"email"`
	Phone     *string `json:"phone" valid:"numeric"`
	Avatar    *string `json:"avatar"`
}
