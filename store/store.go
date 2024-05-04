package store

import (
	"github.com/elbiseu/flashcards/arangodb"
	"github.com/elbiseu/flashcards/interfaces"
)

var (
	Store interfaces.Operator
)

func init() {
	Store = arangodb.NewOperator()
}
