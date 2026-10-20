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

	r.Run(":8080")
}
