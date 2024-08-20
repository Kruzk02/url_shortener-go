package repository

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/go-sql-driver/mysql"
)

var (
	db     *sql.DB
	dbonce sync.Once
)

func InitDB() {
	dbonce.Do(func() {
		var err error
		cfg := mysql.Config{
			User:   os.Getenv("DB_USER"),
			Passwd: os.Getenv("DB_PASSWORD"),
			Net:    "tcp",
			Addr:   fmt.Sprintf("%s:%s", os.Getenv("DB_HOST"), os.Getenv("DB_PORT")),
			DBName: os.Getenv("DB_NAME"),
		}

		db, err = sql.Open("mysql", cfg.FormatDSN())
		if err != nil {
			log.Fatalf("Failed to open database: %v", err)
		}

		if err := db.Ping(); err != nil {
			log.Fatalf("Failed to ping database: %v", err)
		}
		log.Println("Database Connected!")
	})
}

func GetDB() *sql.DB {
	if db == nil {
		log.Fatal("Database connection is not initialized")
	}
	return db
}
