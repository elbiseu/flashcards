package formatters

import (
	"encoding/json"
	"github.com/elbiseu/flashcards/interfaces"
	"github.com/elbiseu/flashcards/structures"
)

func JSON(dto interfaces.DTO) ([]byte, structures.ContentType, error) {
	b, err := json.Marshal(dto.Transferable())
	return b, structures.JSON, err
}
