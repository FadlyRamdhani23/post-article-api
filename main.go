package main

import (
	"article-api/database"
	"article-api/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()

	r := gin.Default()
	r.POST("/article/", handlers.CreatePost)
	r.GET("/article/:id/:offset", handlers.GetPosts)
	r.GET("/article/:id", handlers.GetPost)
	r.PUT("/article/:id", handlers.UpdatePost)
	r.Run(":8080")
}
