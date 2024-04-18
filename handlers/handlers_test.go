package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/elbiseu/flashcards/apitransfers"
	"github.com/elbiseu/flashcards/interfaces"
	"github.com/elbiseu/flashcards/store"
	"github.com/elbiseu/flashcards/structures"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Mocking Store
type mockStore struct{}

func (ms *mockStore) Get(ctx context.Context, key string, document interfaces.Gatherer) error {
	switch key {
	case "577877":
		*(document.(*structures.Flashcard)) = structures.Flashcard{
			Key: key,
			Meta: structures.VerbFlashcardMeta{
				BaseForm:       "shine",
				PastParticiple: "shone",
				PastSimple:     "shone",
			},
			Type:  structures.IrregularVerb,
			Value: "shine",
		}
	case "578068":
		return errors.New("error getting data")
	default:
		return nil
	}
	return nil
}

func (ms *mockStore) Save(ctx context.Context, document interfaces.Gatherer) error {
	return nil
}

func TestFlashcardGetHandler(t *testing.T) {
	store.Store = &mockStore{}
	tests := []struct {
		Name       string
		Setup      func(*http.Request)
		StatusCode int
		Output     apitransfers.Flashcard
	}{
		{
			"ValidKey",
			func(r *http.Request) {
				ctx := r.Context()
				ctx = context.WithValue(ctx, "key", "valid")
				*r = *r.WithContext(ctx)
			},
			http.StatusOK,
			apitransfers.Flashcard{
				BaseForm:       "run",
				PastParticiple: "run",
				PastSimple:     "ran",
				Key:            "valid",
			},
		},
		{
			"InvalidKey",
			func(r *http.Request) {
				ctx := r.Context()
				ctx = context.WithValue(ctx, "key", "invalid")
				*r = *r.WithContext(ctx)
			},
			http.StatusInternalServerError,
			apitransfers.Flashcard{},
		},
		{
			"NoKey",
			func(r *http.Request) {
				*r = *r.WithContext(context.Background())
			},
			http.StatusOK,
			apitransfers.Flashcard{},
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			req, _ := http.NewRequest("GET", "http://localhost:8080/flashcard", nil)
			test.Setup(req)
			w := httptest.NewRecorder()
			flashcardGetHandler(w, req)

			assert.Equal(t, test.StatusCode, w.Result().StatusCode)

			var resp apitransfers.Flashcard
			_ = json.NewDecoder(w.Body).Decode(&resp)

			assert.Equal(t, test.Output, resp)
		})
	}
}
