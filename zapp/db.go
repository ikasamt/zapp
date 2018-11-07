package zapp

import (
	"log"

	"github.com/gin-gonic/gin"
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

// db 接続を返す
func DB(c *gin.Context) (db *gorm.DB) {
	return c.MustGet("DB").(*gorm.DB)
}
