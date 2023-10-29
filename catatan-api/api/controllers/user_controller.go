package controllers

import (
	"catatan-api/api/models"
	"catatan-api/db"
	"net/http"

	"github.com/labstack/echo/v4"
)

// CreateUser creates a new user.
func CreateUser(c echo.Context) error {
	user := new(models.User)
	if err := c.Bind(user); err != nil {
		response := models.BaseResponse{
			Status:  false,
			Message: "Failed create in database",
			Data:    nil,
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	db.DB.Create(&user)

	response := models.BaseResponse{
		Status:  true,
		Message: "User created successfully",
		Data:    user,
	}
	return c.JSON(http.StatusCreated, response)
}

func GetAllUser(c echo.Context) error {
	var users []models.User

	result := db.DB.Preload("Notes").Find(&users)
	if result.Error != nil {
		response := models.BaseResponse{
			Status:  false,
			Message: result.Error.Error(),
			Data:    nil,
		}
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := models.BaseResponse{
		Status:  true,
		Message: "Users retrieved successfully",
		Data:    users,
	}

	return c.JSON(http.StatusOK, response)
}

// GetUser retrieves a user by their ID.
func GetUser(c echo.Context) error {
	userID := c.Param("id")
	var user models.User

	if err := db.DB.First(&user, userID).Error; err != nil {
		response := models.BaseResponse{
			Status:  false,
			Message: "User not found",
			Data:    nil,
		}
		return c.JSON(http.StatusNotFound, response)
	}

	response := models.BaseResponse{
		Status:  true,
		Message: "User retrieved successfully",
		Data:    user,
	}
	return c.JSON(http.StatusOK, response)
}

// UpdateUser updates an existing user.
func UpdateUser(c echo.Context) error {
	userID := c.Param("id")
	var user models.User

	if err := db.DB.First(&user, userID).Error; err != nil {
		response := models.BaseResponse{
			Status:  false,
			Message: "User not found",
			Data:    nil,
		}
		return c.JSON(http.StatusNotFound, response)
	}

	if err := c.Bind(&user); err != nil {
		response := models.BaseResponse{
			Status:  false,
			Message: "Invalid request payload",
			Data:    nil,
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	db.DB.Save(&user)

	response := models.BaseResponse{
		Status:  true,
		Message: "User updated successfully",
		Data:    user,
	}
	return c.JSON(http.StatusOK, response)
}

// DeleteUser deletes a user by their ID.
func DeleteUser(c echo.Context) error {
	userID := c.Param("id")
	var user models.User

	if err := db.DB.First(&user, userID).Error; err != nil {
		response := models.BaseResponse{
			Status:  false,
			Message: "User not found",
			Data:    nil,
		}
		return c.JSON(http.StatusNotFound, response)
	}

	db.DB.Delete(&user)

	response := models.BaseResponse{
		Status:  true,
		Message: "User deleted successfully",
		Data:    nil,
	}
	return c.JSON(http.StatusOK, response)
}
