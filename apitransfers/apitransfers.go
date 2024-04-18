package apitransfers

import (
	"encoding/json"
	"github.com/elbiseu/flashcards/structures"
)

type Flashcard struct {
	BaseForm       string `json:"base_form"`
	Key            string `json:"key"`
	PastParticiple string `json:"past_participle"`
	PastSimple     string `json:"past_simple"`
	Type           string `json:"type"`
	Value          string `json:"value"`
}

func (v *Flashcard) Marshalled() []byte {
	b, _ := json.Marshal(v)
	return b
}

func (v *Flashcard) ContentType() structures.ContentType {
	return structures.JSON
}
