package utils

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDB() {
	var err error
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME)
	fmt.Println(connStr)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error opening database: %q", err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatalf("Error pinging database: %q", err)
	}
	log.Println("Connected to database")
}

func GetDB() *sql.DB {
	if db == nil || db.Ping() != nil {
		InitDB()
	}
	return db
}
