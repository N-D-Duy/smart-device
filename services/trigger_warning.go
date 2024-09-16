package services

import (
	"bytes"
	"health-management/models"
	"net/http"
	"os"
	"strconv"
)

func TriggerNotify(uid string) {
	record := models.HealthRecord{
		UID:            1,
		HeartRate:      200,
		Steps:          1000,
		CaloriesBurned: 500,
		SleepHours:     8.0,
	}

	notify(record)
}

func notify(record models.HealthRecord) {
	//test
	OPENAI_SERVICE_URL := os.Getenv("OPENAI_SERVICE_URL")

	body := []byte(`{
		"uid": ` + strconv.Itoa(int(record.UID)) + `,
		"heart_rate": ` + strconv.Itoa(record.HeartRate) + `,
		"steps": ` + strconv.Itoa(record.Steps) + `,
		"calories_burned": ` + strconv.Itoa(record.CaloriesBurned) + `,
		"sleep_hours": ` + strconv.FormatFloat(record.SleepHours, 'f', -1, 64) + `
	}`)
	//send request to openai service
	http.Post(OPENAI_SERVICE_URL, "application/json", bytes.NewBuffer(body))
}
