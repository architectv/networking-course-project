package models

type Task struct {
	Id        int        `json:"_id,omitempty"`
	ListId    int        `json:"listId"`
	Title     string     `json:"title"`
	Datetimes *Datetimes `json:"datetimes,omitempty"`
	Position  int        `json:"position"`
}

type UpdateTask struct {
	Title     *string          `json:"title"`
	Datetimes *UpdateDatetimes `json:"datetimes,omitempty"`
	NewListId *int             `json:"newListId"`
	Position  *int             `json:"position"`
}
