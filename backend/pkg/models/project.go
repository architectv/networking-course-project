package models

type Project struct {
	Id string `json:"_id,omitempty" bson:"_id,omitempty"`
	OwnerId string `json:"ownerId,omitempty" bson:"ownerId,omitempty"`
	ChatId string `json:"chatId,omitempty" bson:"chatId,omitempty"`
	DefaultPermissions *Permission `json:"defaultPermissions,omitempty" bson:"defaultPermissions,omitempty"`
	Datetimes *Datetimes `json:"datetimes,omitempty" bson:"datetimes,omitempty"`
	Title string `json:"title" binding:"required" bson:"title"`
	Description string `json:"description,omitempty" bson:"description,omitempty"`
}

type ProjectUser struct {
	Id string `json:"_id,omitempty" bson:"_id,omitempty"`
	UserId string `json:"userId,omitempty" bson:"userId,omitempty"`
	ProjectId string `json:"projectId,omitempty" bson:"projectId,omitempty"`
	Permissions *Permission `json:"permissions,omitempty" bson:"permissions,omitempty"`
}
