package middlewares

import (
	"github.com/elbiseu/flashcards/responses"
	"github.com/elbiseu/flashcards/senders"
	"log"
	"net/http"
)

func Middleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				if err = senders.SendResponse(w, responses.InternalServerError); err != nil {
					log.Println(err)
				}
			}
		}()
		next(w, r)
	}
}
