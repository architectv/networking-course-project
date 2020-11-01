package models

type Project struct {
	Id int64 `json:"id,omitempty"`
	OwnerId int64 `json:"ownerId,omitempty"`
	ChatId int64 `json:"chatId,omitempty"`
	DefaultPermissions *Permission `json:"defaultPermissions,omitempty"`
	Datetimes *Datetimes `json:"datetimes,omitempty"`
	Title string `json:"title"`
	Description string `json:"description,omitempty"`
}
