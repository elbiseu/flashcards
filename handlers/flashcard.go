package handlers

import (
	"encoding/json"
	"github.com/elbiseu/flashcards/dtos"
	"github.com/elbiseu/flashcards/entities"
	"github.com/elbiseu/flashcards/instances"
	"github.com/elbiseu/flashcards/serializers"
	"net/http"
)

func Flashcard(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		key := r.PathValue("key")
		flashcardEntity := &entities.Flashcard{
			Key: key,
		}
		if err := instances.Store.Get(flashcardEntity); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		flashcardDTO := &dtos.Flashcard{
			Key:   flashcardEntity.Key,
			Type:  flashcardEntity.Type,
			Value: flashcardEntity.Value,
		}
		b, err := serializers.NewDefaultSerializer(flashcardDTO).Serialize()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if _, err := w.Write(b); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	case http.MethodPost:
		flashcardDTO := &dtos.Flashcard{}
		if err := json.NewDecoder(r.Body).Decode(flashcardDTO); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		flashcardEntity := &entities.Flashcard{
			Key:   flashcardDTO.Key,
			Type:  flashcardDTO.Type,
			Value: flashcardDTO.Value,
		}
		if err := instances.Store.Put(flashcardEntity); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
	default:
		w.WriteHeader(http.StatusNotImplemented)
	}
}
