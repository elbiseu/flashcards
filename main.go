package main

import (
	"context"
	"fmt"
	"github.com/elbiseu/flashcards/arangodb"
	"github.com/elbiseu/flashcards/structures"
	"log"
	"net/http"
	"os"
)

func main() {
	arangoDBUsername, _ := os.LookupEnv("ARANGODB_USERNAME")
	arangoDBPassword, _ := os.LookupEnv("ARANGODB_PASSWORD")
	arangoDBEndpoint, _ := os.LookupEnv("ARANGODB_ENDPOINT")
	arangoDBName, _ := os.LookupEnv("ARANGODB_DB")
	arangoDB, err := arangodb.NewArangoDB(arangoDBUsername, arangoDBPassword, []string{arangoDBEndpoint})
	if err != nil {
		return
	}
	ctx := context.Background()
	database, err := arangoDB.Client.Database(ctx, arangoDBName)
	if err != nil {
		return
	}
	v := structures.Flashcard{
		Value: "",
		Type:  structures.IrregularVerb,
		Meta: structures.Meta{
			BaseForm:       "",
			PastSimple:     "",
			PastParticiple: "",
		},
	}
	if err := arangodb.NewOperation(database).Save(ctx, v); err != nil {
		log.Println(err)
		return
	}
	serveMux := http.NewServeMux()
	serveMux.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./static"))))
	serveMux.HandleFunc("/flashcard/:type", func(responseWriter http.ResponseWriter, request *http.Request) {
		v := structures.Flashcard{
			Value: "",
			Type:  structures.IrregularVerb,
			Meta: structures.Meta{
				BaseForm:       "",
				PastSimple:     "",
				PastParticiple: "",
			},
		}
		fmt.Println(v)
	})
	for e := structures.List.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
	if err := http.ListenAndServe(":8080", serveMux); err != nil {
		log.Fatal(err)
	}
}
