package belajar_golang_database

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestExecSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "INSERT INTO customers(name) VALUES('Husein'), ('Hasan')"
	_, err := db.ExecContext(ctx, script)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new customer")
}

func TestExecSqlParameter(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "INSERT INTO customers(name) VALUES($1), ($2)"
	_, err := db.ExecContext(ctx, script, "Husein", "Hasan")
	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new customer")
}

func TestQuerySql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "SELECT * FROM customers"
	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string

		err := rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}
		fmt.Println("Id:", id)
		fmt.Println("Name:", name)
	}
}

func TestQuerySqlComplex(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := `SELECT id, name, email, balance, rating, "createdAt", "birthDate", married FROM customers`
	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	fmt.Println("=======================")
	for rows.Next() {
		var id, balance int
		var name, email string
		var rating float32
		var createdAt time.Time
		var birthDate sql.NullTime
		var married bool

		err := rows.Scan(&id, &name, &email, &balance, &rating, &createdAt, &birthDate, &married)
		if err != nil {
			panic(err)
		}
		fmt.Println("Id:", id)
		fmt.Println("Name:", name)
		fmt.Println("Email:", email)
		fmt.Println("Balance:", balance)
		fmt.Println("Rating:", rating)
		fmt.Println("Created At:", createdAt)
		if birthDate.Valid {
			fmt.Println("Birth Date:", birthDate.Time)
		} else {
			fmt.Println("Birth Date: NULL")
		}
		fmt.Println("Married:", married)
		fmt.Println("=======================")
	}
}

func TestSqlInjection(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "admin'; --" // mengcomment baris code selanjutnya, yaitu password
	password := "salah"      // padahal password salah tapi tetap berhasil login

	script := fmt.Sprintf("SELECT username FROM users u WHERE u.username = '%s' AND u.\"password\" = '%s' LIMIT 1;", username, password)
	fmt.Println(script)
	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if rows.Next() {
		fmt.Println("Sukses Login!", username)
	} else {
		fmt.Println("Gagal Login!")
	}
}

func TestSqlInjectionSafe(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "admin'; --" // mengcomment baris code selanjutnya, yaitu password
	password := "salah"

	script := "SELECT username FROM users WHERE username = $1 AND password = $2 LIMIT 1;"
	rows, err := db.QueryContext(ctx, script, username, password)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if rows.Next() {
		fmt.Println("Sukses Login!", username)
	} else {
		fmt.Println("Gagal Login!")
	}
}

func TestAutoIncrement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	email := "eko@mail.com"
	comment := "Tes komen"

	script := "INSERT INTO comments(email, comment) VALUES($1, $2) RETURNING id"

	var lastID int64
	err := db.QueryRowContext(ctx, script, email, comment).Scan(&lastID)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new comment with id", lastID)
}

func TestPrepareStatement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "INSERT INTO comments(email, comment) VALUES($1, $2) RETURNING id"
	statement, err := db.PrepareContext(ctx, script)
	if err != nil {
		t.Fatal(err)
	}
	defer statement.Close()

	for i := 0; i < 10; i++ {
		email := "eko" + strconv.Itoa(i) + "@gmail.com"
		comment := "Komentar ke " + strconv.Itoa(i)
		var lastId int64

		err := statement.QueryRowContext(ctx, email, comment).Scan(&lastId)
		if err != nil {
			t.Fatal(err)
		}

		fmt.Println("Comment Id:", lastId)
	}
}

func TestTransaction(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	tx, err := db.Begin()
	if err != nil {
		t.Fatal(err)
	}

	script := "INSERT INTO comments(email, comment) VALUES($1, $2) RETURNING id"
	stmt, err := tx.PrepareContext(ctx, script)
	if err != nil {
		t.Fatal(err)
	}
	defer stmt.Close()

	for i := 0; i < 10; i++ {
		email := "husein" + strconv.Itoa(i) + "@mail.com"
		comment := "Comment-" + strconv.Itoa(i)

		var lastId int64
		err := stmt.QueryRowContext(ctx, email, comment).Scan(&lastId)
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println("Success insert comment dengan id:", lastId)
	}

	err = tx.Commit()
	if err != nil {
		t.Fatal(err)
	}
}
