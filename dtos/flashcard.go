package dtos

import "github.com/elbiseu/flashcards/types"

type Flashcard struct {
	BaseForm       string
	Key            string
	PastParticiple string
	PastSimple     string
	Type           string
	Value          string
}

func (f *Flashcard) Transferable() types.Transferable {
	return map[string]any{
		"base_form":       f.BaseForm,
		"key":             f.Key,
		"past_participle": f.PastParticiple,
		"past_simple":     f.PastSimple,
		"type":            f.Type,
		"value":           f.Value,
	}
}
