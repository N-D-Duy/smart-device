package services

import (
	"health-management/config"
	"health-management/models"
)

// CreateHealthRecordService handles the logic of creating a health record
func CreateHealthRecordService(record models.HealthRecord) error {
	result := config.DB.Exec(
		"INSERT INTO health_records (uid, heart_rate, steps, calories_burned, sleep_hours, created_at, updated_at) VALUES (?, ?, ?, ?, ?, NOW(), NOW())",
		record.UID, record.HeartRate, record.Steps, record.CaloriesBurned, record.SleepHours,
	)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// GetHealthRecordsByUserIDService fetches records based on user ID
func GetHealthRecordsByUserIDService(uid string) ([]models.HealthRecord, error) {

	var records []models.HealthRecord

	rows, err := config.DB.Raw("SELECT * FROM health_records WHERE uid = ?", uid).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var record models.HealthRecord
		if err := rows.Scan(&record.ID, &record.CreatedAt, &record.UpdatedAt, &record.DeletedAt, &record.UID, &record.HeartRate, &record.Steps, &record.CaloriesBurned, &record.SleepHours); err != nil {
			return nil, err
		}
		records = append(records, record)
	}

	return records, nil
}
