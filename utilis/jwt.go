package utilis

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("secret-key")

func CreateToken(username string) (string, error) {

	//Create a new token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	//sign the token with the secret key
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// verify jwt Tokens
func VerifyJWTTokens(tokenString string, w http.ResponseWriter, r *http.Request) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if DisplayErrors(w, "Error parsing jwt token ", err, http.StatusInternalServerError) {
		return
	}

	if DisplayBoolErrors(w, "invalid token", token.Valid, http.StatusInternalServerError) {
		return
	}
}
