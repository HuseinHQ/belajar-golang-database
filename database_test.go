package belajar_golang_database

import (
	"database/sql"
	"fmt"
	"log"
	"testing"

	_ "github.com/lib/pq"
)

func TestOpenConnection(t *testing.T) {
	connStr := "user=postgres dbname=sandbox_db sslmode=verify-full"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(db)
	defer db.Close()
}
