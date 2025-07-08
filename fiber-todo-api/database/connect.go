package database

import (
	"fiber-todo-api/config"
	"fiber-todo-api/models"
	"log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB


func ConnectDB() {
	db, err := gorm.Open(postgres.Open(config.GetDSN()), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Không thể kết nối với DB", err)
	}

    log.Println("✅ Kết nối thành công")
	DB = db

    db.AutoMigrate(&models.Task{})
}