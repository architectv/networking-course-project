package models

type Task struct {
	Id int64 `json:"id,omitempty"`
	BoardId int64 `json:"boardId"`
	ListId int64 `json:"listId"`
	ChatId int64 `json:"chatId,omitempty"`
	Title string `json:"title"`
	Datetimes *Datetimes `json:"datetimes,omitempty"`
}
