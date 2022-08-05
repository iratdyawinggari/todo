package main

import (
	"todo-app/database"
	"todo-app/entity"
	"todo-app/routes"
)

func main() {

	db := database.SetupDB()
	db.AutoMigrate(&entity.ActivityGroups{}, &entity.TodoItems{})

	r := routes.SetupRoutes(db)
	r.Run()
}
