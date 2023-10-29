package main

import (
	"catatan-api/api/routes"
	"catatan-api/db"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// Initialize the database connection
	db.InitDB()

	// Setup routes
	routes.SetupUserRoutes(e)
	routes.SetupNoteRoutes(e)

	// Start the server
	e.Logger.Fatal(e.Start(":8080"))
}
