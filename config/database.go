// config/database.go
package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	envErr := godotenv.Load()
	if envErr != nil {
		log.Fatalf("Error loading .env file")
	}

	// Connection string (chỉnh sửa theo thông tin của bạn)
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")
	host := os.Getenv("POSTGRES_HOST")
	fmt.Println(user, password, dbname)
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable", host, user, password, dbname)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	// Sử dụng connection pooling để tối ưu kết nối
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalf("Failed to get database connection pool: %v", err)
	}

	// Cấu hình connection pooling
	sqlDB.SetMaxOpenConns(25)                 // Số lượng kết nối mở tối đa
	sqlDB.SetMaxIdleConns(25)                 // Số kết nối idle tối đa
	sqlDB.SetConnMaxLifetime(5 * time.Minute) // Thời gian sống tối đa của một kết nối

	log.Println("Connected to database")
}
