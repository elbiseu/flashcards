package stores

import (
	"github.com/elbiseu/flashcards/interfaces"
)

type DefaultStore struct {
}

func NewDefaultStore() *DefaultStore {
	return &DefaultStore{}
}

func (ds DefaultStore) Put(entity interfaces.Entity) error {
	return nil
}

func (ds DefaultStore) Get(entity interfaces.Entity) error {
	return nil
}

func (ds DefaultStore) Remove(entity interfaces.Entity) error {
	return nil
}
