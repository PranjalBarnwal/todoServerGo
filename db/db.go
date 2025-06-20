package db

import (
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var DB *sqlx.DB

func Init() {
	connStr := os.Getenv("DATABASE_URL")
	if connStr == "" {
		log.Fatal("DATABASE_URL is not set")
	}

	var err error
	DB, err = sqlx.Connect("postgres", connStr)
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}

	log.Println("Connected to the database ✅")
}
