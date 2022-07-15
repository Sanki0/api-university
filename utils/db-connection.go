package utils

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ChkError(err error) {
	if err != nil {
		panic(err)
	}
}


func InitDB() {
	db, err := sql.Open("mysql", "test_user:secret@tcp(db:3306)/test_database")
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(25)
	db.SetMaxOpenConns(25)
	db.SetConnMaxLifetime(5*time.Minute)
	db.SetConnMaxIdleTime(3*time.Minute)


	DB = db;
	
}
