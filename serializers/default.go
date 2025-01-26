package serializers

import (
	"encoding/json"
	"github.com/elbiseu/flashcards/interfaces"
	"net/http"
)

type DefaultSerializer struct {
	w http.ResponseWriter
}

func NewDefaultSerializer(w http.ResponseWriter) *DefaultSerializer {
	return &DefaultSerializer{w: w}
}

func (ds *DefaultSerializer) Serialize(dto interfaces.DTO) ([]byte, error) {
	b, err := json.Marshal(dto.ToMap())
	return b, err
}
