package main

import (
	"context"
	"fmt"
	"github.com/elbiseu/flashcards/arangodb"
	"github.com/elbiseu/flashcards/handlers"
	"github.com/elbiseu/flashcards/interfaces"
	"github.com/elbiseu/flashcards/middlewares"
	"net/http"
	"os"
)

var (
	arangoDBUsername, _ = os.LookupEnv("ARANGODB_USERNAME")
	arangoDBPassword, _ = os.LookupEnv("ARANGODB_PASSWORD")
	arangoDBEndpoint, _ = os.LookupEnv("ARANGODB_ENDPOINT")
	arangoDBName, _     = os.LookupEnv("ARANGODB_DB")
	store               interfaces.Storage
)

func init() {
	arangoDB, err := arangodb.NewArangoDB(arangoDBUsername, arangoDBPassword, []string{arangoDBEndpoint})
	if err != nil {
		return
	}
	ctx := context.Background()
	database, err := arangoDB.Client.Database(ctx, arangoDBName)
	if err != nil {
		return
	}
	var db interfaces.Storage = arangodb.NewOperation(database)
	store = db
}

func main() {
	serveMux := http.NewServeMux()
	serveMux.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./static"))))
	serveMux.HandleFunc("/flashcard/{key}", middlewares.Middleware(handlers.Flashcard))
	if err := http.ListenAndServe(":8080", serveMux); err != nil {
		fmt.Printf("The service needs help getting up and running: %v\n", err)
	}
}
