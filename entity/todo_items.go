package entity

type TodoItems struct {
	ID              int    `json:"id"`
	Title           string `json:"title"`
	IsActive        int    `json:"is_active"`
	Priority        string `json:"priority"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
	ActivityGroupId string `json:"activity_group_id"`
}

type TodoItemsListResponse struct {
	Total int         `json:"total"`
	Limit int         `json:"limit"`
	Skip  int         `json:"skip"`
	Data  []TodoItems `json:data"`
}
