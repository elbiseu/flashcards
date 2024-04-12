package arangodb

import (
	"context"
	"github.com/arangodb/go-driver"
	"reflect"
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

func (o *Operator) Get(ctx context.Context, key string, document any) error {
	collection, err := o.getCollection(ctx, document)
	if err != nil {
		return err
	}
	_, err = collection.ReadDocument(ctx, key, &document)
	return err
}

func (o *Operator) Save(ctx context.Context, document any) error {
	collection, err := o.getCollection(ctx, document)
	if err != nil {
		return err
	}
	_, err = collection.CreateDocument(ctx, document)
	return err
}

func (o *Operator) getCollection(ctx context.Context, document any) (driver.Collection, error) {
	collectionName := o.getCollectionName(document)
	collection, err := o.database.Collection(ctx, strings.ToLower(collectionName))
	return collection, err
}

func (o *Operator) getCollectionName(document any) string {
	return reflect.TypeOf(document).Name()
}
