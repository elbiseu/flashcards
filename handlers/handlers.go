package handlers

import (
	"encoding/json"
	"github.com/elbiseu/flashcards/dtos"
	"github.com/elbiseu/flashcards/fixedvalues"
	"github.com/elbiseu/flashcards/formatters"
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
		// TODO: Send response.
	}
}

func flashcardGetHandler(responseWriter http.ResponseWriter, request *http.Request) {
	key := request.PathValue("key")
	var flashcard structures.Flashcard
	if err := store.Store.Get(request.Context(), key, &flashcard); err != nil {
		// TODO: Send response.
		return
	}
	meta, ok := flashcard.Meta.(structures.VerbMeta)
	if !ok {
		// TODO: Send response.
		return
	}
	dto := dtos.Flashcard{
		BaseForm:       meta.BaseForm,
		Key:            flashcard.Key,
		PastParticiple: meta.PastParticiple,
		PastSimple:     meta.PastSimple,
		Type:           flashcard.Type.String(),
		Value:          flashcard.Value,
	}
	senders.NewDefaultSender(responseWriter).SendResponse(&dto, formatters.JSON)
}

func flashcardPostHandler(responseWriter http.ResponseWriter, request *http.Request) {
	var dto dtos.Flashcard
	if err := json.NewDecoder(request.Body).Decode(&dto); err != nil {
		// TODO: Send response.
		return
	}
	flashcard := structures.Flashcard{
		Meta: structures.VerbMeta{
			BaseForm:       dto.BaseForm,
			PastParticiple: dto.PastParticiple,
			PastSimple:     dto.PastSimple,
		},
		Type:  fixedvalues.IrregularVerb,
		Value: dto.Value,
	}
	if err := store.Store.Save(request.Context(), &flashcard); err != nil {
		log.Println(err)
		return
	}
}
