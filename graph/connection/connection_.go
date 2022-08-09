package connection

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func FetchConnection() *sql.DB {
	dsn := "test_user:secret@tcp(db:3306)/test_database"

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(25)
	db.SetMaxOpenConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)
	db.SetConnMaxIdleTime(3 * time.Minute)
	return db
}
