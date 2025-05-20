package db

func Migrate() {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL UNIQUE,
		password TEXT
	);`

	tokenTable := `
	CREATE TABLE IF NOT EXISTS tokens (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER NOT NULL,
		token TEXT NOT NULL UNIQUE,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (user_id) REFERENCES users(id)
	);`

	oauthTokensTable := `
	CREATE TABLE IF NOT EXISTS oauth_tokens (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER NOT NULL,
		provider TEXT NOT NULL,
		access_token TEXT NOT NULL,
		refresh_token TEXT,
		expiry DATETIME,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (user_id) REFERENCES users(id)
	);
	`

	_, err := DB.Exec(oauthTokensTable)
	if err != nil {
		panic("Failed to create oauth_tokens table: " + err.Error())
	}

	_, err = DB.Exec(query)
	if err != nil {
		panic("Failed to run migration: " + err.Error())
	}

	_, err = DB.Exec(tokenTable)
	if err != nil {
		panic("Failed to create tokens table: " + err.Error())
	}
}