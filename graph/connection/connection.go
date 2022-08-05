package connection

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func FetchConnection() *gorm.DB {
	dsn := "test_user:secret@tcp(db)/test_database"

	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	db.DB().SetMaxIdleConns(25)
	db.DB().SetMaxOpenConns(25)
	db.DB().SetConnMaxLifetime(5 * time.Minute)
	db.DB().SetConnMaxIdleTime(3 * time.Minute)
	return db
}
