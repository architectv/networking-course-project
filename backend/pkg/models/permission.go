package models

type Permission struct {
	Read  bool `json:"read,omitempty"`
	Write bool `json:"write,omitempty"`
	Admin bool `json:"admin,omitempty"`
}

type UpdatePermission struct {
	Read  *bool `json:"read,omitempty"`
	Write *bool `json:"write,omitempty"`
	Admin *bool `json:"admin,omitempty"`
}
