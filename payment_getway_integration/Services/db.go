package services

import (
	"fmt"
	"hello/models"

	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {

	const DNS = "host=localhost user=postgres password=password dbname=project port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	database, err := gorm.Open(postgres.Open(DNS), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	database.AutoMigrate(models.Paymentdatails{}, models.TransactionDetails{})

	DB = database
	fmt.Println("Database connection successfully opened")
}
