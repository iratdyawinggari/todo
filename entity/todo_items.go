package entity

import "time"

type TodoItems struct {
	ID              int       `json:"id"`
	Title           string    `json:"title"`
	IsActive        int       `json:"is_active"`
	Priority        string    `json:"priority"`
	Comment         string    `json:"_comment"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	ActivityGroupId int    `json:"activity_group_id"`
}

type TodoItemsListResponse struct {
	Total int         `json:"total"`
	Limit int         `json:"limit"`
	Skip  int         `json:"skip"`
	Data  []TodoItems `json:data"`
}

type UpdateTodotemsInput struct {
	Title    string `json:"title"`
	IsActive int    `json:"is_active"`
	Priority string `json:"priority"`
	Comment  string `json:"_comment"`
}
