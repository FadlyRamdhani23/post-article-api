package main

import (
	"article-api/database"
	"article-api/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()

	r := gin.Default()
	r.Use(cors.Default())
	r.POST("/article/", handlers.CreatePost)
	r.GET("/article/:id/:offset", handlers.GetPosts)
	r.GET("/article/:id", handlers.GetPost)
	r.PUT("/article/:id", handlers.UpdatePost)
	r.DELETE("/article/:id", handlers.DeletePost)
	r.Run(":8080")
}
