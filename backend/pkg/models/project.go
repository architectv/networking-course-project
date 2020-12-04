package models

type Project struct {
	Id                 int         `json:"id,omitempty"`
	OwnerId            int         `json:"ownerId,omitempty"`
	DefaultPermissions *Permission `json:"defaultPermissions,omitempty"`
	Datetimes          *Datetimes  `json:"datetimes,omitempty"`
	Title              string      `json:"title" valid:"length(1|50)"`
	Description        string      `json:"description,omitempty"`
}

type UpdateProject struct {
	DefaultPermissions *UpdatePermission `json:"defaultPermissions,omitempty"`
	Datetimes          *UpdateDatetimes  `json:"datetimes,omitempty"`
	Title              *string           `json:"title" valid:"length(1|50)"`
	Description        *string           `json:"description,omitempty"`
}
