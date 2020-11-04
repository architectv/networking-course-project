package models

type Board struct {
	Id string `json:"_id,omitempty"`
	Title string `json:"title"`
	ProjectId int64 `json:"projectId"`
	DefaultPermissions *Permission `json:"defaultPermissions,omitempty"`
	Datetimes *Datetimes `json:"datetimes,omitempty"`
}
