package users

import (
	"net/http"

	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"

	//custom imports from file
	database "go-task-manager/handlers/database"
	utilis "go-task-manager/utilis"
)

func LoginUser(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		//read request payload
		loginUser, err := utilis.ReadRequestBody(r)
		if utilis.DisplayErrors(w, "Failed to read logging users payload", err, http.StatusInternalServerError) {
			return
		}

		//define variables for request payload
		loginCredentials := utilis.ExtractString(loginUser, "login_credentials")
		password := utilis.ExtractString(loginUser, "password")

		//fetch the stored password hash
		passwordHash, err := database.FetchPasswordHash(db, loginCredentials)
		if utilis.DisplayErrors(w, "Wrong username or email credentials provided", err, http.StatusBadRequest) {
			return
		}

		//verify password hash against the password submitted
		verifyPassword := verifyPassword(password, passwordHash)
		if utilis.DisplayBoolErrors(w, "Password did not match the stored password hash", verifyPassword, http.StatusBadRequest) {
			return
		}

		//generate a new jwt token
		token, err := utilis.CreateToken(loginCredentials)
		if utilis.DisplayErrors(w, "Failed to create a new token", err, http.StatusInternalServerError) {
			return
		}

		//set the token in the response header
		w.Header().Set("Authorization", "Bearer "+token)

		//Define response payload
		response := map[string]interface{}{
			"access_token": token,
			"message":      "user was successfully logged in",
		}

		//write success response
		utilis.CreateSuccessResponse(w, response, http.StatusOK)

	}
}

// verify that given password matches the stored hash
func verifyPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
