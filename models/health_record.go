package models

import (
	"gorm.io/gorm"
)

type HealthRecord struct {
	gorm.Model
	UID            uint `gorm:"index"`
	HeartRate      int
	Steps          int
	CaloriesBurned int
	SleepHours     float64
}
