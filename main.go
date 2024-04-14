package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/elbiseu/flashcards/arangodb"
	"github.com/elbiseu/flashcards/interfaces"
	"github.com/elbiseu/flashcards/responses"
	"github.com/elbiseu/flashcards/structures"
	"github.com/elbiseu/flashcards/transfers"
	"log"
	"net/http"
	"os"
)

var (
	arangoDBUsername, _ = os.LookupEnv("ARANGODB_USERNAME")
	arangoDBPassword, _ = os.LookupEnv("ARANGODB_PASSWORD")
	arangoDBEndpoint, _ = os.LookupEnv("ARANGODB_ENDPOINT")
	arangoDBName, _     = os.LookupEnv("ARANGODB_DB")
	store               interfaces.Store
)

func sendResponse(writer http.ResponseWriter, payload interfaces.APILayer) error {
	writer.Header().Set("Content-Type", "application/json")
	payloadJSON, err := payload.JSON()
	if err != nil {
		return err
	}
	_, err = writer.Write(payloadJSON)
	return err
}

func middleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				if err = sendResponse(w, responses.InternalServerError); err != nil {
					log.Println(err)
				}
			}
		}()
		next(w, r)
	}
}

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
	var db interfaces.Store = arangodb.NewOperation(database)
	store = db
}

func main() {
	serveMux := http.NewServeMux()
	serveMux.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./static"))))
	serveMux.HandleFunc("/flashcard", middleware(func(responseWriter http.ResponseWriter, request *http.Request) {
		switch request.Method {
		case http.MethodPost:
			var verb transfers.Verb
			if err := json.NewDecoder(request.Body).Decode(&verb); err != nil {
				_ = sendResponse(responseWriter, responses.InternalServerError)
			}
			flashcard := structures.Flashcard{
				Value: verb.Value,
				Type:  structures.IrregularVerb,
				Meta: structures.Meta{
					BaseForm:       verb.BaseForm,
					PastSimple:     verb.PastSimple,
					PastParticiple: verb.PastParticiple,
				},
			}
			if err := store.Save(request.Context(), flashcard); err != nil {
				log.Println(err)
				return
			}
		}
	}))
	serveMux.HandleFunc("/flashcard/{key}", middleware(func(responseWriter http.ResponseWriter, request *http.Request) {
		switch request.Method {
		case http.MethodGet:
			key := request.PathValue("key")
			var flashcard structures.Flashcard
			if err := store.Get(request.Context(), key, flashcard); err != nil {
				log.Println(err)
				return
			}
			verb := transfers.Verb{
				Value:          "",
				Type:           "",
				BaseForm:       "",
				PastSimple:     "",
				PastParticiple: "",
			}
			_ = sendResponse(responseWriter, &verb)
		}
	}))
	for e := structures.List.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
	if err := http.ListenAndServe(":8080", serveMux); err != nil {
		log.Fatal(err)
	}
}
