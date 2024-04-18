package apimodifiers

import (
	"github.com/elbiseu/flashcards/structures"
)

type APIModifier struct {
	apiTransfer any
	modifiers   []structures.ModifyingFunc
}

func NewAPIModifier(apiTransfer any) *APIModifier {
	return &APIModifier{apiTransfer: apiTransfer}
}

func (a *APIModifier) Add(modifiers structures.ModifyingFunc) *APIModifier {
	a.modifiers = append(a.modifiers, modifiers)
	return a
}

func (a *APIModifier) Apply() *any {
	for _, modifier := range a.modifiers {
		modifier(&a.apiTransfer)
	}
	return &a.apiTransfer
}
