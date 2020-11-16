package models

type Board struct {
	Id                 string      `json:"_id,omitempty" bson:"_id,omitempty"`
	Title              string      `json:"title" bson:"title"`
	ProjectId          string      `json:"projectId" bson:"projectId"`
	DefaultPermissions *Permission `json:"defaultPermissions,omitempty" bson:"defaultPermissions,omitempty"`
	Datetimes          *Datetimes  `json:"datetimes,omitempty" bson:"datetimes,omitempty"`
}
