package interfaces

import (
	"context"
)

type Store interface {
	Get(ctx context.Context, key string, item any) error
	Save(ctx context.Context, item any) error
}
