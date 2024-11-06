package users

import (
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
		registerUser, err := utilis.ReadRequestBody(r)
		if utilis.DisplayErrors(w, "Failed to read registered users payload", err, http.StatusInternalServerError) {
			return
		}

		//define individual request body
		username := utilis.ExtractString(registerUser, "username")
		email := utilis.ExtractString(registerUser, "email")
		password := utilis.ExtractString(registerUser, "password")
		passwordConfirmation := utilis.ExtractString(registerUser, "password_confirmation")

		//validate that password and password confirmation are the same
		passwordConfirmCheck := confirmPassword(password, passwordConfirmation)
		if utilis.DisplayBoolErrors(w, "Password and confirm password did not match", passwordConfirmCheck, http.StatusBadRequest) {
			return
		}

		//hash password provided
		hashedPassword, err := hashPassword(password)
		if utilis.DisplayErrors(w, "Failed to hash password", err, http.StatusInternalServerError) {
			return
		}

		//insert record into users table
		userId, err := database.InsertToUsers(db, username, email, hashedPassword)
		if utilis.DisplayErrors(w, "Failed to insert record into the users table", err, http.StatusInternalServerError) {
			return
		}

		//Define response payload
		response := map[string]interface{}{
			"id":      userId,
			"message": "User was registered successfully",
		}

		//write success response
		utilis.CreateSuccessResponse(w, response, http.StatusCreated)

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
