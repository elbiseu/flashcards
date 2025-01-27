package middlewares

import (
	"net/http"
)

func DefaultMiddleware(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				http.Error(w, "Internal server error", http.StatusInternalServerError)
				return
			}
		}()
		handler(w, r)
	}
}
