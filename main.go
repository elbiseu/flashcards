package main

import (
	"fmt"
	"github.com/elbiseu/flashcards/handlers"
	"github.com/elbiseu/flashcards/middlewares"
	"net/http"
)

func main() {
	serveMux := http.NewServeMux()
	serveMux.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./static"))))
	serveMux.HandleFunc("/flashcard/{key}", middlewares.Default(handlers.Flashcard))
	if err := http.ListenAndServe(":8080", serveMux); err != nil {
		fmt.Printf("The service needs help getting up and running: %v\n", err)
	}
}
