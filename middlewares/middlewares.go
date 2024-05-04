package middlewares

import (
	"net/http"
)

func Default(handler http.HandlerFunc) http.HandlerFunc {
	return func(responseWriter http.ResponseWriter, request *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				// TODO: Send response.
			}
		}()
		handler(responseWriter, request)
	}
}
