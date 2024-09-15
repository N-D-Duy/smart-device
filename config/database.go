// config/database.go
package config

import (
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	// Connection string (chỉnh sửa theo thông tin của bạn)
	dsn := "user=postgres password=12332145 dbname=test host=db port=5432 sslmode=disable"

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
