package controllers

import (
	"database/sql"
	"net/http"
	"strings"
	// "encoding/hex"
	// "crypto/rand"

	"golang.org/x/crypto/bcrypt"
	"github.com/labstack/echo/v4"

	"github.com/Ech0n/go-oauth-server/db"
	"github.com/Ech0n/go-oauth-server/types"
	"github.com/Ech0n/go-oauth-server/tokens"
	
	_ "github.com/mattn/go-sqlite3"
)



func Register(c echo.Context) error {
	var req types.RegisterRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	req.Email = strings.TrimSpace(req.Email)
	req.Password = strings.TrimSpace(req.Password)

	if req.Email == "" || req.Password == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Email and password required"})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to hash password"})
	}

	query := `INSERT INTO users (email, password) VALUES (?, ?)`
	_, err = db.DB.Exec(query, req.Email, string(hashedPassword))
	if err != nil {
		if strings.Contains(err.Error(), "UNIQUE") {
			return c.JSON(http.StatusConflict, map[string]string{"error": "Email already registered"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Database error"})
	}

	return c.JSON(http.StatusCreated, map[string]string{"message": "User registered successfully"})

}

func Login(c echo.Context) error {
	var req types.LoginRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	req.Email = strings.TrimSpace(req.Email)
	req.Password = strings.TrimSpace(req.Password)

	if req.Email == "" || req.Password == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Email and password are required"})
	}

	var storedHashedPassword string
	var userID int
	err := db.DB.QueryRow("SELECT id, password FROM users WHERE email = ?", req.Email).
		Scan(&userID, &storedHashedPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid email or password"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Database error"})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(storedHashedPassword), []byte(req.Password)); err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid email or password"})
	}

	token, errStr := tokens.RegisterAndGetNewToken(userID);
	if errStr != "" {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": errStr})
	}

	return c.JSON(http.StatusOK, echo.Map{"token": token})
}

func Logout(c echo.Context) error {
	authHeader := c.Request().Header.Get("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Missing or invalid Authorization header"})
	}

	token := strings.TrimPrefix(authHeader, "Bearer ")

	result, err := db.DB.Exec("DELETE FROM tokens WHERE token = ?", token)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to logout"})
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to logout"})
	}

	if rowsAffected == 0 {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Invalid token"})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "Logged out successfully"})
}
