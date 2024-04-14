package apitransfers

import "github.com/elbiseu/flashcards/structures"

type Flashcard struct {
	BaseForm       string
	Key            string
	PastParticiple string
	PastSimple     string
	Type           string
	Value          string
}

func (v *Flashcard) SharedFields() structures.SharedFields {
	return map[string]any{
		"base_form":       v.BaseForm,
		"past_participle": v.PastParticiple,
		"past_simple":     v.PastSimple,
		"type":            v.Type,
		"value":           v.Value,
	}
}

func (v *Flashcard) ContentType() structures.ContentType {
	return structures.JSON
}
