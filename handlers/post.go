package handlers

import (
	"article-api/database"
	"article-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreatePost(c *gin.Context) {

	var input models.Post

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "format JSON tidak valid"})
		return
	}

	query := " INSERT INTO posts (title, content, category, status) VALUES (?, ?, ?, ?)"
	_, err := database.DB.Exec(query, input.Title, input.Content, input.Category, input.Status)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "gagal membuat post"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{})
}
