package database

import (
	"log"

	"github.com/cngJo/golang-api-auth/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Instance *gorm.DB
var dbError error

func Connect(connectionString string) {
	Instance, dbError = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if dbError != nil {
		log.Fatal(dbError)
		panic("failed to connect database")
	}
	log.Println("Database connected")
}

func Migrate() {
	Instance.AutoMigrate(&models.User{})
	log.Println("Database migrated")
}
