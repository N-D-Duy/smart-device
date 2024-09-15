package controllers

import (
	"health-management/config"
	"health-management/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateHealthRecord(c *gin.Context) {
	var record models.HealthRecord
	if err := c.ShouldBindJSON(&record); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := config.DB.Exec(
		"INSERT INTO health_records (uid, heart_rate, steps, calories_burned, sleep_hours, created_at, updated_at) VALUES (?, ?, ?, ?, ?, NOW(), NOW())",
		record.UID, record.HeartRate, record.Steps, record.CaloriesBurned, record.SleepHours,
	)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusCreated, record)
}

func GetHealthRecordsByUserID(c *gin.Context) {
	uid := c.Param("uid")
	var records []models.HealthRecord

	rows, err := config.DB.Raw("SELECT * FROM health_records WHERE uid = ?", uid).Rows()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var record models.HealthRecord
		if err := rows.Scan(&record.ID, &record.CreatedAt, &record.UpdatedAt, &record.DeletedAt, &record.UID, &record.HeartRate, &record.Steps, &record.CaloriesBurned, &record.SleepHours); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		records = append(records, record)
	}
	c.JSON(http.StatusOK, records)
}
