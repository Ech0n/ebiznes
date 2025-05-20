package controllers

import (
	"database/sql"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/Ech0n/go-oauth-server/db"
	_ "github.com/mattn/go-sqlite3"
)

func GetPrivateData(c echo.Context) error {
	authHeader := c.Request().Header.Get("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Missing or invalid Authorization header"})
	}
	token := strings.TrimPrefix(authHeader, "Bearer ")
	var userID int
	err := db.DB.QueryRow("SELECT user_id FROM tokens WHERE token = ?", token).Scan(&userID)
	if (err != nil) {
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Invalid token"})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Database error"})
	}
	return c.JSON(http.StatusOK, echo.Map{"data": "some private data"})
}

