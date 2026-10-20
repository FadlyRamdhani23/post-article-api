package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Connect() {
	dsn := "root:@tcp(127.0.0.1:3306)/post_article?parseTime=true&clientFoundRows=true"

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
