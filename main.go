package main

import (
	"fmt"
	"github.com/elbiseu/flashcards/fileloader"
	"github.com/elbiseu/flashcards/flashcard"
	"github.com/elbiseu/flashcards/verb"
	"log"
	"net/http"
)

func init() {
	file, err := fileloader.NewFileLoader("irregular_verbs.csv")
	if err != nil {
		panic(err)
	}
	records, err := file.CSV()
	if err != nil {
		panic(err)
	}
	for _, record := range records {
		v := flashcard.Flashcard{
			Value: record[0],
			Type:  flashcard.IrregularVerb,
			Meta: verb.Meta{
				BaseForm:       record[0],
				PastSimple:     record[1],
				PastParticiple: record[2],
			},
		}
		verb.List.PushBack(v)
	}
}

func main() {
	serveMux := http.NewServeMux()
	serveMux.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./static"))))
	serveMux.HandleFunc("/flashcard/:type", func(responseWriter http.ResponseWriter, request *http.Request) {})
	for e := verb.List.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
	if err := http.ListenAndServe(":8080", serveMux); err != nil {
		log.Fatal(err)
	}
}
