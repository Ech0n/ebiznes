package tokens

import (
	"encoding/hex"
	"crypto/rand"

	"github.com/Ech0n/go-oauth-server/db"
	_ "github.com/mattn/go-sqlite3"
)

func RegisterAndGetNewToken(userID int) (string, string) {
	token, err := generateToken()
	if err != nil {
		return "", "Internal server error"
	}

	_, err = db.DB.Exec(`INSERT INTO tokens (user_id, token) VALUES (?, ?)`, userID, token)
	if err != nil {
		return "", "Internal server error"
	}
	return token, ""
}

func generateToken() (string, error) {
	bytes := make([]byte, 32)
	_, err := rand.Read(bytes)
	return hex.EncodeToString(bytes), err
}

