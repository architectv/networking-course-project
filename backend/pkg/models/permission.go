package models

type Permission struct {
	Read  bool `json:"read,omitempty"`
	Write bool `json:"write,omitempty"`
	Grant bool `json:"grant,omitempty"`
}
