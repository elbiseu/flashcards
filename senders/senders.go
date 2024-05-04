package senders

import (
	"github.com/elbiseu/flashcards/interfaces"
	"github.com/elbiseu/flashcards/types"
	"log"
	"net/http"
)

type DefaultSender struct {
	responseWriter http.ResponseWriter
}

func NewDefaultSender(responseWriter http.ResponseWriter) *DefaultSender {
	return &DefaultSender{responseWriter: responseWriter}
}

func (ds *DefaultSender) SendResponse(dto interfaces.DTO, converterFunc types.ConverterFunc) {
	b, contentType, err := converterFunc(dto)
	if err != nil {
		log.Println(err)
		return
	}
	ds.responseWriter.Header().Set("Content-Type", contentType.String())
	if _, err := ds.responseWriter.Write(b); err != nil {
		log.Println(err)
		return
	}
}
