package middlewares

import (
	"net/http"

	//import custom libraries
	utilis "go-task-manager/utilis"
)

type verifyToken struct{}

func (m *verifyToken) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	token := r.Header.Get("access_token")
	if token != "" {
		utilis.VerifyJWTTokens(token, w, r)
	}
	next(w, r)
}

func VerifyToken() *verifyToken {
	return &verifyToken{}
}
