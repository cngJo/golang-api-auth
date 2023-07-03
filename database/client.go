package database

import (
	"log"

	"github.com/cngJo/golang-api-auth/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var Instance *gorm.DB
var dbError error

func Connect(connectionString string) {
	Instance, dbError = gorm.Open(mysql.Open(connectionString), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if dbError != nil {
		log.Fatal(dbError)
		panic("failed to connect database")
	}
	log.Println("Database connected")
}

func Migrate() {
	Instance.AutoMigrate(&models.User{})
	Instance.AutoMigrate(&models.RefreshToken{})

	log.Println("Database migrated")
}
