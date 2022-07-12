package utils

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func ChkError(err error) {
	if err != nil {
		panic(err)
	}
}

func PingDb(db *sql.DB) {
	err := db.Ping()
	ChkError(err)
}

func ConnectionDB() *sql.DB {
	db, err := sql.Open("mysql", "test_user:secret@tcp(db:3306)/test_database")
	if err != nil {
		panic(err.Error())
	}
	return db
}
