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
	r.POST("/activity-groups", controllers.CreateActivityGroups)
	r.GET("/activity-groups/:id", controllers.GetActivityGroupsById)
	r.PATCH("/activity-groups/:id", controllers.UpdateActivityGroups)
	r.DELETE("activity-groups/:id", controllers.DeleteActivityGroups)

	r.GET("/todo-items", controllers.GetAllTodoItems)
	r.POST("/todo-items", controllers.CreateTodoItems)
	r.GET("/todo-items/:id", controllers.GetTodoItemsById)
	r.PATCH("/todo-items/:id", controllers.UpdateTodoItems)
	r.DELETE("/todo-items/:id", controllers.DeleteTodoItems)

	return r
}
