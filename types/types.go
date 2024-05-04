package types

import (
	"github.com/elbiseu/flashcards/interfaces"
)

type ConverterFunc func(dto interfaces.DTO) ([]byte, ContentType, error)
type Transferable map[string]any
