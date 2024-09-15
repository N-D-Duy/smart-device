package routes

import (
	"health-management/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	// Routes cho User controller
	r.GET("/users", controllers.GetUsers)
	r.POST("/users", controllers.CreateUser)

	// Routes cho HealthRecord controller
	r.POST("/health-records", controllers.CreateHealthRecord)
	r.GET("/health-records/:uid", controllers.GetHealthRecordsByUserID)
}
