package db

import (
	"log"

	"github.com/Adilfarooque/todolist/utils/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
var DB *gorm.DB // Global DB variable to store the connection

func InitDB() {
    dsn := "host=localhost user=postgres password=7356 dbname=todolist port=5432 sslmode=disable"
    var err error
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{}) // Assign to global DB variable

    if err != nil {
        log.Fatal("Failed to connect to the database:", err)
    }

    if err = DB.AutoMigrate(&models.Todo{}); err != nil {
        log.Fatal("Failed to migrate database:", err)
    }
}
