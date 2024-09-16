package services

import (
	"health-management/models"
	"log"
	"time"
)

func ScheduleGen(uid uint) {
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	records := []models.HealthRecord{}
	var recordCount int

	for {
		select {
		case <-ticker.C:
			// Generate mock health data every minute
			newRecord := MockGen(uid)
			records = append(records, newRecord)
			recordCount++

			log.Printf("Generated record for user %d: %+v\n", uid, newRecord)

			// After 1 hour (60 records), save to database
			if recordCount >= 60 {
				CreateHealthRecordService(newRecord)
				records = []models.HealthRecord{} // Reset the record slice for the next hour
				recordCount = 0
			}
		}
	}
}

func ScheduledGenStart(uid uint) {
	go ScheduleGen(uid)
}
