package handlers

import (
	"article-api/database"
	"article-api/models"
	"database/sql"
	"net/http"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/gin-gonic/gin"
)

func CreatePost(c *gin.Context) {

	var input models.Post

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "format JSON tidak valid"})
		return
	}
	if errs := validatePost(input); len(errs) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"errors": errs})
		return
	}

	query := " INSERT INTO posts (title, content, category, status) VALUES (?, ?, ?, ?)"
	_, err := database.DB.Exec(query, input.Title, input.Content, input.Category, input.Status)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "gagal membuat post"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "article berhasil dibuat"})
}

func GetPosts(c *gin.Context) {

	limit, err1 := strconv.Atoi(c.Param("id"))
	offset, err2 := strconv.Atoi(c.Param("offset"))

	if err1 != nil || err2 != nil || limit < 0 || offset < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "limit dan offset harus angka"})
		return
	}

	query := "SELECT id, title, content, category, created_date, updated_date, status FROM posts LIMIT ? OFFSET ?"
	rows, err := database.DB.Query(query, limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "gagal mengambil data"})
		return
	}
	defer rows.Close()

	posts := []models.Post{}
	for rows.Next() {
		var p models.Post
		err := rows.Scan(&p.ID, &p.Title, &p.Content, &p.Category, &p.CreatedDate, &p.UpdatedDate, &p.Status)
		if err = rows.Err(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "gagal membaca data"})
			return
		}

		posts = append(posts, p)
	}
	c.JSON(http.StatusOK, posts)
}

func GetPost(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id harus angka"})
		return
	}

	var p models.Post
	query := "SELECT id, title, content, category, created_date, updated_date, status FROM posts WHERE id = ?"
	err = database.DB.QueryRow(query, id).Scan(&p.ID, &p.Title, &p.Content, &p.Category, &p.CreatedDate, &p.UpdatedDate, &p.Status)

	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "article tidak ditemukan"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "gagal mengambil data"})
		return
	}
	c.JSON(http.StatusOK, p)
}

func UpdatePost(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id harus angka"})
		return
	}

	var input models.Post
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "format JSON tidak valid"})
		return
	}
	if errs := validatePost(input); len(errs) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"errors": errs})
		return
	}

	query := "UPDATE posts SET title = ?, content = ?, category = ?, status = ? WHERE id = ?"
	result, err := database.DB.Exec(query, input.Title, input.Content, input.Category, input.Status, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "gagal memperbarui post"})
		return
	}

	affected, _ := result.RowsAffected()
	if affected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "article tidak ditemukan"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "article berhasil diperbarui"})
}

func DeletePost(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id harus angka"})
		return
	}
	result, err := database.DB.Exec("DELETE FROM posts WHERE id = ?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "gagal menghapus article"})
		return
	}
	affected, _ := result.RowsAffected()
	if affected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "article tidak ditemukan"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "article berhasil dihapus"})
}

func validatePost(p models.Post) map[string]string {
	errs := map[string]string{}

	if utf8.RuneCountInString(strings.TrimSpace(p.Title)) < 20 {
		errs["title"] = "title wajib diisi, minimal 20 karakter"
	}
	if utf8.RuneCountInString(strings.TrimSpace(p.Content)) < 200 {
		errs["content"] = "content wajib diisi, minimal 200 karakter"
	}
	if utf8.RuneCountInString(strings.TrimSpace(p.Category)) < 3 {
		errs["category"] = "category wajib diisi, minimal 3 karakter"
	}
	if p.Status != "publish" && p.Status != "draft" && p.Status != "thrash" {
		errs["status"] = "status harus publish, draft, atau thrash"
	}

	return errs
}
