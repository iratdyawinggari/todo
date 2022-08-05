package controllers

import (
	"net/http"
	"time"
	"todo-app/entity"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

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

func CreateTodoItems(c *gin.Context) {
	var input entity.TodoItems
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

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

func GetTodoItemsById(c *gin.Context) {
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

func UpdateTodoItems(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)

	var existingTodoItems entity.TodoItems
	if err := db.Where("id = ?", c.Param("id")).First(&existingTodoItems).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

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
	db := c.MustGet("db").(*gorm.DB)
	var todoItems entity.TodoItems
	if err := db.Where("id = ?", c.Param("id")).First(&todoItems).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&todoItems)

	c.JSON(http.StatusOK, nil)
}
