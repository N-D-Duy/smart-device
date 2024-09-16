package controllers

import (
	"health-management/models"
	"health-management/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetUsersController handles the HTTP request to fetch all users
func GetUsersController(c *gin.Context) {
	users, err := services.GetUsersService()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

// CreateUserController handles the HTTP request to create a new user
func CreateUserController(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := services.CreateUserService(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully", "user": user})
}
