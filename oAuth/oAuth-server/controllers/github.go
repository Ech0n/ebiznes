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
	"golang.org/x/oauth2/github"
	"github.com/Ech0n/go-oauth-server/db"
	"github.com/Ech0n/go-oauth-server/tokens"
	"github.com/joho/godotenv"
)

var githubOauthConfig *oauth2.Config

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

    githubOauthConfig = &oauth2.Config{
        ClientID:     os.Getenv("GITHUB_CLIENT_ID"),
        ClientSecret: os.Getenv("GITHUB_CLIENT_SECRET"),
        RedirectURL:  "http://localhost:8080/github/callback",
        Scopes:       []string{"user:email"},
        Endpoint:     github.Endpoint,
    }
}

func GitHubLogin(c echo.Context) error {
    url := githubOauthConfig.AuthCodeURL("random-state")
    return c.Redirect(http.StatusTemporaryRedirect, url)
}

func GitHubCallback(c echo.Context) error {
    code := c.QueryParam("code")

    token, err := githubOauthConfig.Exchange(context.Background(), code)
    if err != nil {
        return c.String(http.StatusInternalServerError, "Error exchanging code for token")
    }

    client := githubOauthConfig.Client(context.Background(), token)

    resp, err := client.Get("https://api.github.com/user/emails")
    if err != nil {
        return c.String(http.StatusInternalServerError, "Error fetching user emails")
    }
    defer resp.Body.Close()

    var emails []struct {
        Email   string `json:"email"`
        Primary bool   `json:"primary"`
        Verified bool  `json:"verified"`
    }
    if err := json.NewDecoder(resp.Body).Decode(&emails); err != nil {
        return c.String(http.StatusInternalServerError, "Error decoding emails")
    }

    var userEmail string
    for _, e := range emails {
        if e.Primary && e.Verified {
            userEmail = e.Email
            break
        }
    }
    if userEmail == "" && len(emails) > 0 {
        userEmail = emails[0].Email
    }

    if userEmail == "" {
        return c.String(http.StatusInternalServerError, "No email found")
    }

    userID, err := db.FindOrCreateOAuthUserByEmail(userEmail)
    if err != nil {
        return c.String(http.StatusInternalServerError, "Error creating/finding user")
    }

    tokenString, errMsg := tokens.RegisterAndGetNewToken(userID)
    if errMsg != "" {
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

    redirectURL := fmt.Sprintf("http://localhost:5173/oauth-success?token=%s", tokenString)
    return c.Redirect(http.StatusTemporaryRedirect, redirectURL)
}