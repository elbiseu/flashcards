package handlers

import (
	"encoding/json"
	"github.com/elbiseu/flashcards/dtos"
	"github.com/elbiseu/flashcards/entities"
	"github.com/elbiseu/flashcards/serializers"
	"net/http"
)

func Flashcard(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		// key := request.PathValue("key")
		entity := entities.Flashcard{}
		// TODO: Get entity from database.
		dto := dtos.Flashcard{
			Key:   entity.Key,
			Type:  entity.Type,
			Value: entity.Value,
		}
		b, err := serializers.NewDefaultSerializer(w).Serialize(&dto)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if _, err := w.Write(b); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	case http.MethodPost:
		var dto dtos.Flashcard
		if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
			return
		}
		/*
			flashcard := entities.Flashcard{
				Key:   dto.Key,
				Type:  dto.Type,
				Value: dto.Value,
			}
		*/
		// TODO: Store entity in database.
		w.WriteHeader(http.StatusCreated)
	default:
		// TODO: Send response.
	}
}
