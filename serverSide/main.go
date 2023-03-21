package main

import (
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/nashkispace/ns-todolist/serverSide/routes"
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
	router.DELETE("/api/deleteTask/:id", routes.DeleteTask)
	router.Run(":" + port)
}
