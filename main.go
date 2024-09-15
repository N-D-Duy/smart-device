// main.go
package main

import (
	"health-management/config"
	"health-management/models"
	"health-management/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDatabase()

	err := config.DB.AutoMigrate(&models.User{}, &models.HealthRecord{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	r := gin.Default()
	routes.SetupRoutes(r)

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
