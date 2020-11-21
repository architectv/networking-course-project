package models

type Project struct {
	Id                 int         `json:"id,omitempty"`
	OwnerId            int         `json:"ownerId,omitempty"`
	DefaultPermissions *Permission `json:"defaultPermissions,omitempty"`
	Datetimes          *Datetimes  `json:"datetimes,omitempty"`
	Title              string      `json:"title" valid:"length(0|50)"`
	Description        string      `json:"description,omitempty"`
}
