package zapp

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// GetDBInstance returns DBインスタンスを返す
func GetDBInstance(dsn string) (db *gorm.DB) {
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		log.Println(err)
	}
	db.LogMode(true)
	return db
}
