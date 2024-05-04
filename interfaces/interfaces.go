package interfaces

import (
	"context"
	"github.com/elbiseu/flashcards/types"
)

type APITransfer interface {
	Transferable() ([]byte, types.ContentType, error)
}

type Gatherer interface {
	Gathering() string
	Identify() string
	Database() string
}

type Storage interface {
	Name() string
}

type Operator interface {
	Get(ctx context.Context, gatherer Gatherer) error
	Save(ctx context.Context, gatherer Gatherer) error
}
