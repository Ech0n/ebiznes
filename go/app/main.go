package main

import (
	// "project/controllers"
	"project/database"
	"project/routes"
	"path/filepath"
	"github.com/labstack/echo/v4"
)

func main() {
	println("Starting the application...")
	e := echo.New()

	database.ConnectDB()
	database.Migrate()

	routes.SetupRoutes(e)

	e.GET("/", func(c echo.Context) error {
		viewPath := filepath.Join("views", "index.html")
		return c.File(viewPath)
	})

	e.Logger.Fatal(e.Start(":8080"))
}
