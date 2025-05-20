package db

import (
	"database/sql"
	"log"
	"time"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "./app.db")
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	if err := DB.Ping(); err != nil {
		log.Fatal("Database unreachable:", err)
	}
}

func SaveOAuthToken(userID int, provider, accessToken, refreshToken string, expiry time.Time) error {
	_, err := DB.Exec(`
		INSERT INTO oauth_tokens (user_id, provider, access_token, refresh_token, expiry)
		VALUES (?, ?, ?, ?, ?)
	`, userID, provider, accessToken, refreshToken, expiry)
	return err
}

func FindOrCreateOAuthUserByEmail(email string) (int, error) {
	var userID int
	err := DB.QueryRow(`SELECT id FROM users WHERE email = ?`, email).Scan(&userID)
	if err == sql.ErrNoRows {
		res, err := DB.Exec(`INSERT INTO users (email) VALUES (?)`, email)
		if err != nil {
			return 0, err
		}
		lastID, err := res.LastInsertId()
		return int(lastID), err
	} else if err != nil {
		return 0, err
	}
	return userID, nil
}