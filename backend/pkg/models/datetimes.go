package models

type Datetimes struct {
	Created  int64 `json:"created,omitempty"`
	Updated  int64 `json:"updated,omitempty"`
	Accessed int64 `json:"accessed,omitempty"`
}

type UpdateDatetimes struct {
	Created  *int64 `json:"created,omitempty"`
	Updated  *int64 `json:"updated,omitempty"`
	Accessed *int64 `json:"accessed,omitempty"`
}
