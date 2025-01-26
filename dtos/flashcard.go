package dtos

type Flashcard struct {
	Key   string
	Type  string
	Value string
}

func (f *Flashcard) ToMap() map[string]any {
	return map[string]any{
		"key":   f.Key,
		"type":  f.Type,
		"value": f.Value,
	}
}
