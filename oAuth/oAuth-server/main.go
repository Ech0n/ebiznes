package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/Ech0n/go-oauth-server/routes"
	"github.com/Ech0n/go-oauth-server/db"
)

func main() {
	db.InitDB()
	db.Migrate()
	
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:5173","http://localhost",},
		AllowMethods: []string{
			"GET", "POST", "PUT", "DELETE", "OPTIONS",
		},
		AllowHeaders: []string{"Content-Type", "Authorization", "X-Requested-With", "X-Real-IP",},
	}))

	routes.UseEmailPasswordAuthRoutes(e)
	routes.UseGoogleAuthRoutes(e)
	routes.UseGithubAuthRoutes(e)

	e.Logger.Fatal(e.Start(":8080"))
}
