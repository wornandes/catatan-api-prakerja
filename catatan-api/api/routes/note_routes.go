package routes

import (
	"catatan-api/api/controllers"

	"github.com/labstack/echo/v4"
)

func SetupNoteRoutes(e *echo.Echo) {
	e.POST("/notes", controllers.CreateNote)
	e.GET("/notes", controllers.GetAllNote)
	e.GET("/notes/:id", controllers.GetNote)
	e.PUT("/notes/:id", controllers.UpdateNote)
	e.DELETE("/notes/:id", controllers.DeleteNote)
}
