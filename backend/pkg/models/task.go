package models

type Task struct {
	Id        string     `json:"_id,omitempty"`
	ListId    string     `json:"listId"`
	Title     string     `json:"title"`
	Datetimes *Datetimes `json:"datetimes,omitempty"`
	Position  int        `json:"position"`
}

type UpdateTask struct {
	Title     *string    `json:"title"`
	Datetimes *Datetimes `json:"datetimes,omitempty"`
	Position  *int       `json:"position"`
}
