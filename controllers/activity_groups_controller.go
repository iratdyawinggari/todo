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
func GetAllActivityGroups(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var activityGroups []entity.ActivityGroups
	db.Find(&activityGroups)

	result := entity.ActivityGroupsListResponse{
		Total: len(activityGroups),
		Limit: 1000,
		Skip:  0,
		Data:  activityGroups,
	}

	c.JSON(http.StatusOK, gin.H{"": result})
}

// POST /tasks
// Create new task
func CreateTask(c *gin.Context) {
	// Validate input
	var input entity.CreateActivityGroupsInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create task
	activityGroups := entity.ActivityGroups{
		Title:     input.Title,
		Email:     input.Email,
		Comment:   input.Comment,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	db := c.MustGet("db").(*gorm.DB)
	db.Create(&activityGroups)

	c.JSON(http.StatusOK, gin.H{"data": activityGroups})
}

// GET /tasks/:id
// Find a task
// func FindTask(c *gin.Context) { // Get model if exist
// 	var task models.Task

// 	db := c.MustGet("db").(*gorm.DB)
// 	if err := db.Where("id = ?", c.Param("id")).First(&task).Error; err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"data": task})
// }

// PATCH /tasks/:id
// Update a task
// func UpdateTask(c *gin.Context) {

// 	db := c.MustGet("db").(*gorm.DB)
// 	// Get model if exist
// 	var task models.Task
// 	if err := db.Where("id = ?", c.Param("id")).First(&task).Error; err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
// 		return
// 	}

// 	// Validate input
// 	var input UpdateTaskInput
// 	if err := c.ShouldBindJSON(&input); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	date := "2006-01-02"
// 	deadline, _ := time.Parse(date, input.Deadline)

// 	var updatedInput models.Task
// 	updatedInput.Deadline = deadline
// 	updatedInput.AssingedTo = input.AssingedTo
// 	updatedInput.Task = input.Task

// 	db.Model(&task).Updates(updatedInput)

// 	c.JSON(http.StatusOK, gin.H{"data": task})
// }

// DELETE /tasks/:id
// Delete a task
// func DeleteTask(c *gin.Context) {
// 	// Get model if exist
// 	db := c.MustGet("db").(*gorm.DB)
// 	var book models.Task
// 	if err := db.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
// 		return
// 	}

// 	db.Delete(&book)

// 	c.JSON(http.StatusOK, gin.H{"data": true})
// }
