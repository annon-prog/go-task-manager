package database

import (
	"github.com/jmoiron/sqlx"
)

func FetchPasswordHash(db *sqlx.DB, credentials string) (string, error) {

	var passwordHash string
	query := `SELECT password FROM users WHERE username = $1 OR email = $2  `
	err := db.QueryRow(query, credentials, credentials).Scan(&passwordHash)
	if err != nil {
		return "", err
	}
	return passwordHash, nil

}
