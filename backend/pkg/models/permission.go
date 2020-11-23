package models

type Permission struct {
	Read  bool `json:"read,omitempty" valid:"type(bool)"`
	Write bool `json:"write,omitempty" valid:"type(bool)"`
	Admin bool `json:"admin,omitempty" valid:"type(bool)"`
}

type UpdatePermission struct {
	Read  *bool `json:"read,omitempty" valid:"type(bool)"`
	Write *bool `json:"write,omitempty" valid:"type(bool)"`
	Admin *bool `json:"admin,omitempty" valid:"type(bool)"`
}
