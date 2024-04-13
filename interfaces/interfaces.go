package interfaces

import (
	"context"
)

type APILayer interface {
	JSON() ([]byte, error)
}

type APIError interface {
}

type Store interface {
	Get(ctx context.Context, key string, item any) error
	Save(ctx context.Context, item any) error
}
