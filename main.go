package main

import (
	"article-api/database"
	"fmt"
	"log"
	"net/http"
)

func main() {
	database.Connect()

	query := `
	CREATE TABLE IF NOT EXISTS posts (
		id INT AUTO_INCREMENT PRIMARY KEY,
		title VARCHAR(200),
		content TEXT,
		category VARCHAR(100),
		created_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
		status VARCHAR(100)
	)`

	_, err := database.DB.Exec(query)
	if err != nil {
		log.Fatal("migrate gagal: ", err)
	}

	fmt.Println("migrate selesai")
	log.Println("Table created successfully")
	log.Println("Starting server on :8080")
	http.ListenAndServe(":8080", nil)
}
