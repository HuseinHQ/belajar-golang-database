package belajar_golang_database

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/lib/pq"
)

func GetConnection() *sql.DB {
	connStr := "user=postgres dbname=sandbox_db password=postgres"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(1 * time.Hour)

	return db
}
