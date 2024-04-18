package senders

import (
	"github.com/elbiseu/flashcards/interfaces"
	"log"
	"net/http"
)

type APISender struct {
	responseWriter http.ResponseWriter
}

func NewAPISender(responseWriter http.ResponseWriter) *APISender {
	return &APISender{responseWriter: responseWriter}
}

func (s APISender) SendResponse(apiSender interfaces.APITransfer) {
	s.responseWriter.Header().Set("Content-Type", string(apiSender.ContentType()))
	if _, err := s.responseWriter.Write(apiSender.Marshalled()); err != nil {
		log.Println(err)
	}
}
