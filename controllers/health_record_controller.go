package controllers

import (
	"health-management/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetHealthRecordsByUserIDController handles fetching health records by user ID
func GetHealthRecordsByUserIDController(c *gin.Context) {
	uid := c.Param("uid")
	records, err := services.GetHealthRecordsByUserIDService(uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, records)
}
