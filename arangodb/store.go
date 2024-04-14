package arangodb

import (
	"context"
	"github.com/arangodb/go-driver"
	"github.com/elbiseu/flashcards/interfaces"
	"strings"
)

type Operator struct {
	database driver.Database
}

func NewOperation(database driver.Database) *Operator {
	return &Operator{
		database: database,
	}
}

func (o *Operator) Get(ctx context.Context, key string, document interfaces.Gatherer) error {
	collection, err := o.getCollection(ctx, document)
	if err != nil {
		return err
	}
	_, err = collection.ReadDocument(ctx, key, &document)
	return err
}

func (o *Operator) Save(ctx context.Context, document interfaces.Gatherer) error {
	collection, err := o.getCollection(ctx, document)
	if err != nil {
		return err
	}
	_, err = collection.CreateDocument(ctx, document)
	return err
}

func (o *Operator) getCollection(ctx context.Context, document interfaces.Gatherer) (driver.Collection, error) {
	collection, err := o.database.Collection(ctx, strings.ToLower(document.Gathering()))
	return collection, err
}
