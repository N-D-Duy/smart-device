package routes

import (
	"health-management/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	// Routes cho User controller
	r.GET("/users", controllers.GetUsersController)
	r.POST("/register", controllers.CreateUserController)

	// Routes cho HealthRecord controller
	r.GET("/health-records/:uid", controllers.GetHealthRecordsByUserIDController)
}
