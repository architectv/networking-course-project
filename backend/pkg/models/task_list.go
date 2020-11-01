package models

type TaskList struct {
	Id int64 `json:"id,omitempty"`
	Title string `json:"title"`
	BoardId int64 `json:"boardId"`
	Position int32 `json:"position"`
}
