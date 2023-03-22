package main

import (
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/nashkispace/ns-todolist/serverSide/api"
)

func main() {

	http.HandleFunc("/api/task", func(w http.ResponseWriter, r *http.Request) {
		// CORS
		w.Header().Add("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Add("Access-Control-Allow-Methods", "POST")
		w.Header().Add("Access-Control-Allow-Credentials", "true")

	})

	http.HandleFunc("/api/tasks", func(w http.ResponseWriter, r *http.Request) {
		// CORS
		w.Header().Add("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Add("Access-Control-Allow-Methods", "GET")
		w.Header().Add("Access-Control-Allow-Credentials", "true")

	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "4000"
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(cors.Default())

	router.POST("/api/task", api.AddTask)
	router.GET("/api/tasks", api.GetTasks)
	// router.DELETE("/api/deleteTask/:id", api.DeleteTask)
	router.Run(":" + port)
}
