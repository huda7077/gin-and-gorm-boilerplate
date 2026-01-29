package configs

import (
	"log"
	"os"

	"github.com/huda7077/gin-and-gorm-boilerplate/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := os.Getenv("DATABASE_URL")
	// Example:
	// postgres://user:password@localhost:5432/dbname?sslmode=disable
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database")
	}
	if err := db.AutoMigrate(models.All...); err != nil {
		log.Fatal(err)
	}
	DB = db
	log.Println("Database connected")
}
