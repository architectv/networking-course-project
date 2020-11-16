package models

type Permission struct {
	Read bool `json:"read,omitempty" bson:"read,omitempty"`
	Write bool `json:"write,omitempty" bson:"write,omitempty"`
	Access bool `json:"access,omitempty" bson:"access,omitempty"`
}
