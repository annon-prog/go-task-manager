package middlewares

import (
	"net/http"
)

type fetchToken struct{}

func (f *fetchToken) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	token := r.Header.Get("access_token")
	if token != "" {
		w.Header().Set("Authorization", token)
	}
	next(w, r)
}

func FetchToken() *fetchToken {
	return &fetchToken{}
}
