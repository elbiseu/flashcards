package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/elbiseu/flashcards/arangodb"
	"github.com/elbiseu/flashcards/interfaces"
	"github.com/elbiseu/flashcards/structures"
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

type CustomRes struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func sendResponse(w http.ResponseWriter, status int, message string, data interface{}) {
	res := CustomRes{
		Status:  status,
		Message: message,
		Data:    data,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	// TODO.
	_ = json.NewEncoder(w).Encode(res)
}

func middleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Recovered from panic: %v", err)
				http.Error(w, "", http.StatusInternalServerError)
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
			v := structures.Flashcard{
				Value: "",
				Type:  structures.IrregularVerb,
				Meta: structures.Meta{
					BaseForm:       "",
					PastSimple:     "",
					PastParticiple: "",
				},
			}
			if err := store.Save(request.Context(), v); err != nil {
				log.Println(err)
				return
			}
		}
	}))
	serveMux.HandleFunc("/flashcard/{key}", middleware(func(responseWriter http.ResponseWriter, request *http.Request) {
		switch request.Method {
		case http.MethodGet:
			key := request.PathValue("key")
			var v structures.Flashcard
			if err := store.Get(request.Context(), key, v); err != nil {
				log.Println(err)
				return
			}
			sendResponse(responseWriter, http.StatusOK, "OK", v)
		}
	}))
	for e := structures.List.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
	if err := http.ListenAndServe(":8080", serveMux); err != nil {
		log.Fatal(err)
	}
}
