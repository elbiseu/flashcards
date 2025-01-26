package entities

type Flashcard struct {
	Key   string
	Type  string
	Value string
}

func (f *Flashcard) Container() string {
	return "flashcard"
}
