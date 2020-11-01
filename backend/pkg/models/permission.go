package models

type Permission struct {
	Read bool `json:"read,omitempty"`
	Write bool `json:"write,omitempty"`
	Access bool `json:"access,omitempty"`
}
