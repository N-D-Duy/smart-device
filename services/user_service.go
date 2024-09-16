package services

import (
	"health-management/config"
	"health-management/models"
)

// GetUsersService retrieves all users from the database
func GetUsersService() ([]models.User, error) {
	var users []models.User

	rows, err := config.DB.Raw("SELECT * FROM users").Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt, &user.UID); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

// CreateUserService inserts a new user into the database
func CreateUserService(user models.User) error {
	result := config.DB.Exec("INSERT INTO users (uid, created_at, updated_at) VALUES (?, NOW(), NOW())", user.UID)
	if result.Error != nil {
		return result.Error
	}

	//trigger the scheduled health record generation for the new user
	ScheduledGenStart(user.UID)
	return nil
}
