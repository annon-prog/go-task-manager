package users

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"

	//custom imports from file
	database "go-task-manager/handlers/database"
	utilis "go-task-manager/utilis"
)

func RegisterUser(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		//read request payload
		registeredUser, err := utilis.ReadRequestBody(r)
		if err != nil {
			log.Printf("Failed to read registered users payload : %v", err)
			http.Error(w, "Failed to read registered users payload", http.StatusInternalServerError)
			return
		}

		//define individual request body
		username := utilis.ExtractString(registeredUser, "username")
		email := utilis.ExtractString(registeredUser, "email")
		password := utilis.ExtractString(registeredUser, "password")
		passwordConfirmation := utilis.ExtractString(registeredUser, "password_confirmation")

		//validate that password and password confirmation are the same
		passwordConfirmCheck := confirmPassword(password, passwordConfirmation)
		if !passwordConfirmCheck {
			log.Println("Confirm password is different from the password provided")
			http.Error(w, "Confirm password is different from the password provided", http.StatusBadRequest)
			return
		}

		//hash password provided
		hashedPassword, err := hashPassword(password)
		if err != nil {
			log.Printf("Failed to hash password : %v", err)
			http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		}

		//insert record into users table
		userId, err := database.InsertToUsers(db, username, email, hashedPassword)
		if err != nil {
			log.Printf("Failed to insert record into the users table: %v", err)
			http.Error(w, "Failed to insert record into the users table", http.StatusInternalServerError)
			return
		}

		response := map[string]interface{}{
			"id":      userId,
			"message": "User was registered successfully",
		}

		//write json response
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response)

	}

}

func confirmPassword(password string, passwordConfirmation string) bool {
	if password == passwordConfirmation {
		return true
	} else {
		return false
	}
}

func hashPassword(password string) (string, error) {
	//generates a hash from the password provided
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
