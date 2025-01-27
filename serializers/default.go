package serializers

import (
	"encoding/json"
	"github.com/elbiseu/flashcards/interfaces"
)

type DefaultSerializer struct {
	dto interfaces.DTO
}

func NewDefaultSerializer(dto interfaces.DTO) *DefaultSerializer {
	return &DefaultSerializer{dto: dto}
}

func (ds *DefaultSerializer) Serialize() ([]byte, error) {
	return json.Marshal(ds.dto.ToMap())
}
