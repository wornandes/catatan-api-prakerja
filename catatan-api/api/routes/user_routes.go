package routes

import (
	"catatan-api/api/controllers"

	"github.com/labstack/echo/v4"
)

func SetupUserRoutes(e *echo.Echo) {
	e.POST("/users", controllers.CreateUser)
	e.GET("/users", controllers.GetAllUser)
	e.GET("/users/:id", controllers.GetUser)
	e.PUT("/users/:id", controllers.UpdateUser)
	e.DELETE("/users/:id", controllers.DeleteUser)
}
