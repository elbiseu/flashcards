package arangodb

import (
	"context"
	"github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
	"github.com/elbiseu/flashcards/interfaces"
	"os"
	"strings"
)

var (
	arangoDBUsername, _ = os.LookupEnv("ARANGODB_USERNAME")
	arangoDBPassword, _ = os.LookupEnv("ARANGODB_PASSWORD")
	arangoDBEndpoint, _ = os.LookupEnv("ARANGODB_ENDPOINT")
)

type Operator struct {
	client driver.Client
}

func NewOperator() *Operator {
	connection, err := http.NewConnection(http.ConnectionConfig{
		Endpoints: []string{arangoDBEndpoint},
	})
	if err != nil {
		panic(err)
	}
	client, err := driver.NewClient(driver.ClientConfig{
		Authentication: driver.BasicAuthentication(arangoDBUsername, arangoDBPassword),
		Connection:     connection,
	})
	if err != nil {
		panic(err)
	}
	return &Operator{
		client: client,
	}
}

func (o *Operator) Get(ctx context.Context, gatherer interfaces.Gatherer) error {
	database, err := o.do(ctx, gatherer)
	if err != nil {
		return err
	}
	_, err = collection.ReadDocument(ctx, key, &gatherer)
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

func (o *Operator) do(ctx context.Context, document interfaces.Gatherer) (driver.Collection, error) {
	database, err := o.client.Database(ctx, document.Database())
	if err != nil {
		return nil, err
	}
	collection, err := database.Collection(ctx, strings.ToLower(document.Gathering()))
	return collection, err
}
