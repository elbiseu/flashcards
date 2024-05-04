package interfaces

import "github.com/elbiseu/flashcards/types"

type DTO interface {
	Transferable() types.Transferable
}
