package repository

import (
	"database/sql"
	"log"
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
			User:   "root",
			Passwd: "Password",
			Net:    "tcp",
			Addr:   "127.0.0.1:3306",
			DBName: "url",
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
