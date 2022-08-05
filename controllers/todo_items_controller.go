package controllers

import (
	"net/http"
	"time"
	"todo-app/entity"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// GET /tasks
// Get all tasks
func GetAllTodoItems(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var todoItems []entity.TodoItems
	db.Find(&todoItems)

	result := entity.TodoItemsListResponse{
		Total: len(todoItems),
		Limit: 1000,
		Skip:  0,
		Data:  todoItems,
	}

	c.JSON(http.StatusOK, result)
}

// POST /tasks
// Create new task
func CreateTodoItems(c *gin.Context) {
	// Validate input
	var input entity.TodoItems
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create task
	todoItems := entity.TodoItems{
		Title:           input.Title,
		Comment:         input.Comment,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
		ActivityGroupId: input.ActivityGroupId,
	}

	db := c.MustGet("db").(*gorm.DB)
	db.Create(&todoItems)

	c.JSON(http.StatusOK, todoItems)
}

// GET /tasks/:id
// Find a task
func GetTodoItemsById(c *gin.Context) { // Get model if exist
	var todoItems entity.TodoItems

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&todoItems).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	result := entity.TodoItems{
		ID:       todoItems.ID,
		Title:    todoItems.Title,
		IsActive: todoItems.IsActive,
		Priority: todoItems.Priority,
	}
	c.JSON(http.StatusOK, result)
}

// PATCH /tasks/:id
// Update a task
func UpdateTodoItems(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var existingTodoItems entity.TodoItems
	if err := db.Where("id = ?", c.Param("id")).First(&existingTodoItems).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input entity.UpdateTodotemsInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedInput entity.TodoItems
	updatedInput.Title = input.Title
	updatedInput.IsActive = input.IsActive
	updatedInput.Priority = input.Priority
	updatedInput.Comment = input.Comment
	updatedInput.UpdatedAt = time.Now()

	db.Model(&existingTodoItems).Updates(updatedInput)

	c.JSON(http.StatusOK, existingTodoItems)
}

func DeleteTodoItems(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var todoItems entity.TodoItems
	if err := db.Where("id = ?", c.Param("id")).First(&todoItems).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&todoItems)

	c.JSON(http.StatusOK, nil)
}
