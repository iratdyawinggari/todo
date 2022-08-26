package entity

import "time"

type ActivityGroups struct {
	ID        int       `json:"id" gorm:"primary_key`
	Title     string    `json:"title"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"json:"updated_at"`
	Email     string    `json:"email"`
	Comment   string    `json:"_comment"`
	TodoItem  []TodoItems `gorm:"ForeignKey:ActivityGroupId"` 	
}

type ActivityGroupsDetailResponse struct {
	ID        int         `json:"id"`
	Title     string      `json:"title"`
	CreatedAt time.Time   `json:"created_at"`
	TodoItems []TodoItems `json:todo_items,gorm:"foreignKey:ID;"`
}

type ActivityGroupsListResponse struct {
	Total int              `json:"total"`
	Limit int              `json:"limit"`
	Skip  int              `json:"skip"`
	Data  []ActivityGroups `json:data"`
}

type CreateActivityGroupsInput struct {
	Title   string `json:"title"`
	Email   string `json:"email"`
	Comment string `json:"_comment"`
}

type UpdateActivityGroupsInput struct {
	Title string `json:"title"`
}
