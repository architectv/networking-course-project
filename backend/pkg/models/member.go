package models

type Member struct {
	Id          int         `json:"id"`
	Nickname    string      `json:"nickname" valid:"length(3|32)"`
	Avatar      string      `json:"avatar"`
	IsOwner     bool        `json:"isOwner"`
	Permissions *Permission `json:"permissions"`
}
