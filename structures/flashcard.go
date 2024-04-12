package structures

const (
	IrregularVerb Type = "Irregular Verb"
)

type Type string

type Flashcard struct {
	Value string
	Type  Type
	Meta  any
}
