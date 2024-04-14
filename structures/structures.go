package structures

import "github.com/elbiseu/flashcards/interfaces"

const (
	IrregularVerb FlashcardType = "Irregular Verb"
	JSON          ContentType   = "JSON"
)

type ContentType string
type FlashcardType string
type SharedFields map[string]any
type BBody []byte
type ModifyingFunc func(transfer *interfaces.APITransfer)

type Flashcard struct {
	Key   string
	Meta  any
	Type  FlashcardType
	Value string
}

func (v *Flashcard) Gathering() string {
	return "flashcard"
}

func (v *Flashcard) Identify() string {
	return v.Key
}

type VerbFlashcardMeta struct {
	BaseForm       string
	PastParticiple string
	PastSimple     string
}
