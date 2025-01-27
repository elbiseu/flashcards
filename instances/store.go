package instances

import (
	"github.com/elbiseu/flashcards/interfaces"
	"github.com/elbiseu/flashcards/stores"
)

var Store interfaces.Store

func init() {
	Store = stores.NewDefaultStore()
}
