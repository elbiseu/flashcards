package handlers

import (
	"encoding/json"
	"github.com/elbiseu/flashcards/apimodifiers"
	"github.com/elbiseu/flashcards/apitransfers"
	"github.com/elbiseu/flashcards/interfaces"
	"github.com/elbiseu/flashcards/responses"
	"github.com/elbiseu/flashcards/senders"
	"github.com/elbiseu/flashcards/structures"
	"log"
	"net/http"
)

func f1Test(transfer *interfaces.APITransfer) {
	(*transfer).SharedFields()[""] = 1
}

func Flashcard(responseWriter http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodGet:
		key := request.PathValue("key")
		var verbFlashcard structures.Flashcard
		if err := store.Get(request.Context(), key, verbFlashcard); err != nil {
			log.Println(err)
		}
		meta := verbFlashcard.Meta.(structures.VerbFlashcardMeta)
		transfer := apitransfers.Flashcard{
			BaseForm:       meta.BaseForm,
			Key:            verbFlashcard.Key,
			PastParticiple: meta.PastParticiple,
			PastSimple:     meta.PastSimple,
			Type:           string(verbFlashcard.Type),
			Value:          verbFlashcard.Value,
		}
		apimodifiers.NewAPIModifier(&transfer).Add(f1Test).Add().Apply()
		senders.NewAPISender(responseWriter).SendResponse(&transfer)
	case http.MethodPost:
		var verb apitransfers.Flashcard
		if err := json.NewDecoder(request.Body).Decode(&verb); err != nil {
			_ = senders.SendResponse(responseWriter, responses.InternalServerError)
		}
		flashcard := structures.Flashcard{
			Meta: structures.VerbFlashcardMeta{
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
}
