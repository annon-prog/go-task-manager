package users

import (
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"
)

//function registerUser
//receive request body
//define individual request body payload
//insert record into the users table
//fetch the id from the record
//return the id and a message saying user was registered successfully

func registerUser(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		//read request payload
		registeredUser, err := readBody.readRequestBody(r)
		if err != nil {
			log.Printf("Failed to read registered users payload : %v", err)
			http.Error(w, "Failed to read registered users payload", http.StatusInternalServerError)
			return
		}

		//define individual request body
		username, ok := registeredUser["username"].(string)
		if !ok {
			log.Printf("Failed to read username from the request payload")
		}

	}

}
