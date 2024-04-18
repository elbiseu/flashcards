package interfaces

import (
	"context"
	"github.com/elbiseu/flashcards/structures"
)

type APIModifier interface {
	Apply() APIModifier
}

type APITransfer interface {
	ContentType() structures.ContentType
	Marshalled() []byte
}

type Gatherer interface {
	Gathering() string
	Identify() string
}

type Storage interface {
	Get(ctx context.Context, key string, gatherer Gatherer) error
	Save(ctx context.Context, gatherer Gatherer) error
}