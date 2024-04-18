package store

import (
	"context"
	"github.com/elbiseu/flashcards/arangodb"
	"github.com/elbiseu/flashcards/interfaces"
	"os"
)

var (
	arangoDBUsername, _ = os.LookupEnv("ARANGODB_USERNAME")
	arangoDBPassword, _ = os.LookupEnv("ARANGODB_PASSWORD")
	arangoDBEndpoint, _ = os.LookupEnv("ARANGODB_ENDPOINT")
	arangoDBName, _     = os.LookupEnv("ARANGODB_DB")
	Store               interfaces.Storage
)

func init() {
	arangoDB, err := arangodb.NewArangoDB(arangoDBUsername, arangoDBPassword, []string{arangoDBEndpoint})
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	database, err := arangoDB.Client.Database(ctx, arangoDBName)
	if err != nil {
		panic(err)
	}
	Store = arangodb.NewOperation(database)
}
