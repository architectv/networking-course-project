package models

type Project struct {
	Id                 int         `json:"_id,omitempty"`
	OwnerId            int         `json:"owner_id,omitempty"`
	DefaultPermissions *Permission `json:"defaultPermissions,omitempty"`
	Datetimes          *Datetimes  `json:"datetimes,omitempty"`
	Title              string      `json:"title" valid:"length(0|50)"`
	Description        string      `json:"description,omitempty"`
}
