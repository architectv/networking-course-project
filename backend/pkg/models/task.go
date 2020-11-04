package models

type Task struct {
	Id string `json:"_id,omitempty"`
	BoardId string `json:"boardId"`
	ListId string `json:"listId"`
	ChatId string `json:"chatId,omitempty"`
	Title string `json:"title"`
	Datetimes *Datetimes `json:"datetimes,omitempty"`
}
