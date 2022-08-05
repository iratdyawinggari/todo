package controllers

import (
	"net/http"
	"time"
	"todo-app/entity"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

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

	c.JSON(http.StatusOK, result)
}

func CreateActivityGroups(c *gin.Context) {
	var input entity.CreateActivityGroupsInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	activityGroups := entity.ActivityGroups{
		Title:     input.Title,
		Email:     input.Email,
		Comment:   input.Comment,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	db := c.MustGet("db").(*gorm.DB)
	db.Create(&activityGroups)

	c.JSON(http.StatusOK, activityGroups)
}

func GetActivityGroupsById(c *gin.Context) {
	var activityGroups entity.ActivityGroups

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&activityGroups).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	result := entity.ActivityGroupsDetailResponse{
		ID:        activityGroups.ID,
		Title:     activityGroups.Title,
		CreatedAt: activityGroups.CreatedAt,
		TodoItems: nil,
	}
	c.JSON(http.StatusOK, result)
}

func UpdateActivityGroups(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)

	var existingActivityGroups entity.ActivityGroups
	if err := db.Where("id = ?", c.Param("id")).First(&existingActivityGroups).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var input entity.UpdateActivityGroupsInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedInput entity.ActivityGroups
	updatedInput.Title = input.Title
	updatedInput.UpdatedAt = time.Now()

	db.Model(&existingActivityGroups).Updates(updatedInput)

	c.JSON(http.StatusOK, existingActivityGroups)
}

func DeleteActivityGroups(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var activityGroups entity.ActivityGroups
	if err := db.Where("id = ?", c.Param("id")).First(&activityGroups).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&activityGroups)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
