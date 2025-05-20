package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/Ech0n/go-oauth-server/controllers"
)

func UseEmailPasswordAuthRoutes(e *echo.Echo) {
	authGroup := e.Group("/auth")

	authGroup.POST("/register", controllers.Register)
	authGroup.POST("/login", controllers.Login)
	authGroup.DELETE("/logout", controllers.Logout)

	e.GET("/data", controllers.GetPrivateData)
}

func UseGoogleAuthRoutes(e *echo.Echo) {
	authGroup := e.Group("/google")

	authGroup.GET("/callback", controllers.GoogleCallback)
	authGroup.GET("/login", controllers.GoogleLogin)

}

func UseGithubAuthRoutes(e *echo.Echo) {
	authGroup := e.Group("/github")

	authGroup.GET("/callback", controllers.GitHubCallback)
	authGroup.GET("/login", controllers.GitHubLogin)

}