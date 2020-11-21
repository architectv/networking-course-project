package models

type Board struct {
	Id                 string      `json:"id,omitempty"`
	ProjectId          int         `json:"projectId"`
	DefaultPermissions *Permission `json:"defaultPermissions,omitempty"`
	Datetimes          *Datetimes  `json:"datetimes,omitempty"`
	Title              string      `json:"title"`
}
