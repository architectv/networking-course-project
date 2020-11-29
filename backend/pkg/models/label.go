package models

type Label struct {
	Id      string `json:"id,omitempty"`
	BoardId int    `json:"boardId"`
	Name    string `json:"name"`
	Color   uint32 `json:"color"`
}

type UpdateLabel struct {
	Name  *string `json:"name"`
	Color *uint32 `json:"color"`
}
