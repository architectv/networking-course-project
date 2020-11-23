package models

type Task struct {
	Id        int        `json:"_id,omitempty"`
	ListId    int        `json:"listId"`
	Title     string     `json:"title"`
	Datetimes *Datetimes `json:"datetimes,omitempty"`
	Position  int        `json:"position" valid:"type(int)"`
}

type UpdateTask struct {
	Title     *string          `json:"title"`
	ListId    *int             `json:"listId" valid:"type(*int)"`
	Datetimes *UpdateDatetimes `json:"datetimes,omitempty"`
	Position  *int             `json:"position" valid:"type(*int)"`
}
