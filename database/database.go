package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}

func Connect() {
	user := getEnv("DB_USER", "root")
	pass := getEnv("DB_PASS", "")
	host := getEnv("DB_HOST", "127.0.0.1")
	port := getEnv("DB_PORT", "3306")
	name := getEnv("DB_NAME", "article")

	dsn := user + ":" + pass + "@tcp(" + host + ":" + port + ")/" + name + "?parseTime=true&clientFoundRows=true"

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(" database connect")
	// sql.Open belum benar-benar konek, jadi dicek pakai Ping
	if err = db.Ping(); err != nil {
		log.Fatal("gagal konek ke database: ", err)
	}

	DB = db
}
