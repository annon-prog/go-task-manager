package database

import (
	"log"

	"github.com/jmoiron/sqlx"
)

func InsertToUsers(db *sqlx.DB, username string, email string, password string) (int, error) {

	query := `INSERT INTO users(username, email, password) VALUES (?, ?, ?)`

	result, err := db.Exec(query, username, email, password)
	if err != nil {
		log.Printf("Failed to insert record into users table : %v", err)
		return 0, err
	}
	userId, err := result.LastInsertId()
	if err != nil {
		log.Printf("Failed to retrieve user id : %v", err)
		return 0, err
	}

	return int(userId), nil
}
