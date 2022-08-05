package entity

type ActivityGroups struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Email     string `json:"age"`
}

type ActivityGroupsDetailResponse struct {
	ID        int         `json:"id"`
	Title     string      `json:"title"`
	CreatedAt string      `json:"created_at"`
	TodoItems []TodoItems `json:todo_items,gorm:"foreignKey:ID;"`
}

type ActivityGroupsListResponse struct {
	Total int              `json:"total"`
	Limit int              `json:"limit"`
	Skip  int              `json:"skip"`
	Data  []ActivityGroups `json:data"`
}
