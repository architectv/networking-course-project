package models

type Project struct {
	Id string `json:"_id,omitempty"`
	OwnerId string `json:"ownerId,omitempty"`
	ChatId string `json:"chatId,omitempty"`
	DefaultPermissions *Permission `json:"defaultPermissions,omitempty"`
	Datetimes *Datetimes `json:"datetimes,omitempty"`
	Title string `json:"title"`
	Description string `json:"description,omitempty"`
}
