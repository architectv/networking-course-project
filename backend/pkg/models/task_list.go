package models

type TaskList struct {
	Id string `json:"_id,omitempty"`
	Title    string `json:"title"`
	BoardId int64 `json:"boardId"`
	Position int32 `json:"position"`
}
