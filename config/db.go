package config

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB

func InitDB() {
	var err error
	db, err = gorm.Open("sqlite3", "./gorm.db")
	if err != nill {
		log.Fatal(err)
	}

	db.LogMode(gin.Mode() == gin.DebugMode)
}

func GetDB() *gorm.DB {
	return db
}

func CloseDB() {
	db.Close()
}
