package services

import (
	"health-management/models"
	"math/rand"
)

var (
	heartRateMin       = 60
	heartRateMax       = 200
	stepsMin           = 0
	stepsMax           = 1000
	caloriesMin        = 50
	caloriesMax        = 500
	sleepHoursMin      = 6.0
	sleepHoursMax      = 8.0
	heartRateLast      = rand.Intn(heartRateMax-heartRateMin) + heartRateMin // Initialize with a random starting value
	stepsLast          = rand.Intn(stepsMax-stepsMin) + stepsMin
	caloriesBurnedLast = rand.Intn(caloriesMax-caloriesMin) + caloriesMin
	sleepHoursLast     = sleepHoursMin + (rand.Float64() * (sleepHoursMax - sleepHoursMin))
)

func MockGen(uid uint) models.HealthRecord {
	// Create random small variations to previous values
	heartRateChange := rand.Intn(5) - 2          // Change heart rate by -2 to +2
	stepsChange := rand.Intn(50)                 // Increase steps by 0 to 50
	caloriesChange := rand.Intn(20) - 10         // Change calories burned by -10 to +10
	sleepChange := (rand.Float64() * 0.5) - 0.25 // Change sleep hours by -0.25 to +0.25

	// Ensure values stay within limits and adjust smoothly
	heartRateLast = clamp(heartRateLast+heartRateChange, heartRateMin, heartRateMax)
	stepsLast = clamp(stepsLast+stepsChange, stepsMin, stepsMax)
	caloriesBurnedLast = clamp(caloriesBurnedLast+caloriesChange, caloriesMin, caloriesMax)
	sleepHoursLast = clampFloat(sleepHoursLast+sleepChange, sleepHoursMin, sleepHoursMax)

	return models.HealthRecord{
		UID:            uid,
		HeartRate:      heartRateLast,
		Steps:          stepsLast,
		CaloriesBurned: caloriesBurnedLast,
		SleepHours:     sleepHoursLast,
	}
}

// Helper function to ensure values stay within a certain range
func clamp(value, min, max int) int {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}

func clampFloat(value, min, max float64) float64 {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}
