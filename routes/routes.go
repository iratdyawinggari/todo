package routes

import (
	"todo-app/controllers"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func SetupRoutes(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})
	r.GET("/activity-groups", controllers.GetAllActivityGroups)
	r.POST("/activity-groups", controllers.CreateTask)
	// r.GET("/activity-groups/:id", controllers.FindTask)
	// r.PATCH("/activity-groups/:id", controllers.UpdateTask)
	// r.DELETE("activity-groups/:id", controllers.DeleteTask)
	return r
}
