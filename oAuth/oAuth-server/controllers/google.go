package controllers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"fmt"

	"github.com/labstack/echo/v4"
	// "github.com/labstack/echo/v4/middleware"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"github.com/Ech0n/go-oauth-server/db"
	"github.com/Ech0n/go-oauth-server/tokens"
	"github.com/joho/godotenv"
)

var googleOauthConfig *oauth2.Config

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	googleOauthConfig = &oauth2.Config{
		RedirectURL:  "http://localhost:8080/google/callback",
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}
}

func GoogleLogin(c echo.Context) error {
	url := googleOauthConfig.AuthCodeURL("random-state")
	log.Println("Redirecting to Google OAuth URL:", url)
	return c.Redirect(http.StatusTemporaryRedirect, url)
}



func GoogleCallback(c echo.Context) error {
	
	code := c.QueryParam("code")
	log.Println("Received code from Google:", code)
	token, err := googleOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Error while exchanging code for token")
	}
	log.Println("Access token obtained:", token.AccessToken)


	client := googleOauthConfig.Client(context.Background(), token)

	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		return c.String(http.StatusInternalServerError, "Error fetching user info")
	}
	defer resp.Body.Close()

	var userInfo map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		return c.String(http.StatusInternalServerError, "Error decoding user info")
	}

	email, ok := userInfo["email"].(string)
	if !ok {
		log.Println("Email not found in userInfo:", userInfo)
		return c.String(http.StatusInternalServerError, "Email not found")
	}

	userID, err := db.FindOrCreateOAuthUserByEmail(email)
	if err != nil {
		log.Println("Error finding/creating user:", err)
		return c.String(http.StatusInternalServerError, "User DB error")
	}
	tokenStr, errMsg := tokens.RegisterAndGetNewToken(userID)
	if errMsg != "" {
		log.Println("Token registration error:", errMsg)
		return c.String(http.StatusInternalServerError, errMsg)
	}

	expiry := token.Expiry
	refreshToken := ""
	if token.RefreshToken != "" {
		refreshToken = token.RefreshToken
	}

	err = db.SaveOAuthToken(userID, "google", token.AccessToken, refreshToken, expiry)
	if err != nil {
        return c.String(http.StatusInternalServerError, "Error saving oAuth data to db")
	}


	redirectURL := fmt.Sprintf("http://localhost:5173/oauth-success?token=%s", tokenStr)
	log.Println("Redirecting user to:", redirectURL)
	return c.Redirect(http.StatusTemporaryRedirect, redirectURL)
}