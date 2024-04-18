package handlers

import (
	"encoding/json"
	"github.com/elbiseu/flashcards/apitransfers"
	"github.com/elbiseu/flashcards/senders"
	"github.com/elbiseu/flashcards/store"
	"github.com/elbiseu/flashcards/structures"
	"log"
	"net/http"
)

func Flashcard(responseWriter http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodGet:
		flashcardGetHandler(responseWriter, request)
	case http.MethodPost:
		flashcardPostHandler(responseWriter, request)
	default:
		http.Error(responseWriter, "", http.StatusBadRequest)
	}
}

func flashcardGetHandler(responseWriter http.ResponseWriter, request *http.Request) {
	key := request.PathValue("key")
	var verbFlashcard structures.Flashcard
	if err := store.Store.Get(request.Context(), key, &verbFlashcard); err != nil {
		log.Println(err)
	}
	meta, ok := verbFlashcard.Meta.(structures.VerbFlashcardMeta)
	if !ok {
		return
	}
	transfer := apitransfers.Flashcard{
		BaseForm:       meta.BaseForm,
		Key:            verbFlashcard.Key,
		PastParticiple: meta.PastParticiple,
		PastSimple:     meta.PastSimple,
		Type:           string(verbFlashcard.Type),
		Value:          verbFlashcard.Value,
	}
	senders.NewAPISender(responseWriter).SendResponse(&transfer)
}

func flashcardPostHandler(responseWriter http.ResponseWriter, request *http.Request) {
	var verb apitransfers.Flashcard
	if err := json.NewDecoder(request.Body).Decode(&verb); err != nil {
		// TODO: Send errors.
		// senders.NewAPISender(responseWriter).SendResponse()
	}
	flashcard := structures.Flashcard{
		Meta: structures.VerbFlashcardMeta{
			BaseForm:       verb.BaseForm,
			PastSimple:     verb.PastSimple,
			PastParticiple: verb.PastParticiple,
		},
	}
	if err := store.Store.Save(request.Context(), &flashcard); err != nil {
		log.Println(err)
		return
	}
}
