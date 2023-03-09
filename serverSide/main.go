package main

import (
	"os"

	"/ns-todolist/serverSide/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "4000"
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(cors.Default())

	router.POST("/api/task", routes.AddTask)
	router.GET("/api/tasks", routes.GetTasks)
	// router.GET("/entry/:id/", routes.GetEntryById)
	// router.GET("/ingredient/:ingredient", routes.GetEntriesByIngredient)

	// router.PUT("/entry/update/:id", routes.UpdateEntry)
	// router.PUT("/ingredient/update/:id", routes.UpdateIngredient)
	router.DELETE("/api/deleteTask/:id", routes.DeleteTask)
	router.Run(":" + port)
}
