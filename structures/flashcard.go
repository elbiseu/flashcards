package structures

import "github.com/elbiseu/flashcards/types"

type Flashcard struct {
	Key   string
	Meta  any
	Type  types.FlashcardType
	Value string
}

func (v *Flashcard) Gathering() string {
	return "flashcard"
}

func (v *Flashcard) Identify() string {
	return v.Key
}
