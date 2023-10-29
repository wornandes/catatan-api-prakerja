package controllers

import (
	"catatan-api/api/models"
	"catatan-api/db"
	"net/http"

	"github.com/labstack/echo/v4"
)

// CreateNote creates a new note.
func CreateNote(c echo.Context) error {
	note := new(models.Note)
	if err := c.Bind(note); err != nil {
		response := models.BaseResponse{
			Status:  false,
			Message: "Failed create in database",
			Data:    nil,
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	db.DB.Create(&note)

	response := models.BaseResponse{
		Status:  true,
		Message: "Note created successfully",
		Data:    note,
	}
	return c.JSON(http.StatusCreated, response)
}

func GetAllNote(c echo.Context) error {
	var notes []models.Note
	result := db.DB.Find(&notes)
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
		Message: "Notes retrieved successfully",
		Data:    notes,
	}

	return c.JSON(http.StatusOK, response)
}

// GetNote retrieves a note by its ID.
func GetNote(c echo.Context) error {
	noteID := c.Param("id")
	var note models.Note

	if err := db.DB.First(&note, noteID).Error; err != nil {
		response := models.BaseResponse{
			Status:  false,
			Message: "Note not found",
			Data:    nil,
		}
		return c.JSON(http.StatusNotFound, response)
	}

	response := models.BaseResponse{
		Status:  true,
		Message: "Note retrieved successfully",
		Data:    note,
	}
	return c.JSON(http.StatusOK, response)
}

// UpdateNote updates an existing note.
func UpdateNote(c echo.Context) error {
	noteID := c.Param("id")
	var note models.Note

	if err := db.DB.First(&note, noteID).Error; err != nil {
		response := models.BaseResponse{
			Status:  false,
			Message: "Note not found",
			Data:    nil,
		}
		return c.JSON(http.StatusNotFound, response)
	}

	if err := c.Bind(&note); err != nil {
		response := models.BaseResponse{
			Status:  false,
			Message: "Invalid request payload",
			Data:    nil,
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	db.DB.Save(&note)

	response := models.BaseResponse{
		Status:  true,
		Message: "Note updated successfully",
		Data:    note,
	}
	return c.JSON(http.StatusOK, response)
}

// DeleteNote deletes a note by its ID.
func DeleteNote(c echo.Context) error {
	noteID := c.Param("id")
	var note models.Note

	if err := db.DB.First(&note, noteID).Error; err != nil {
		response := models.BaseResponse{
			Status:  false,
			Message: "Note not found",
			Data:    nil,
		}
		return c.JSON(http.StatusNotFound, response)
	}

	db.DB.Delete(&note)

	response := models.BaseResponse{
		Status:  true,
		Message: "Note deleted successfully",
		Data:    nil,
	}
	return c.JSON(http.StatusOK, response)
}
