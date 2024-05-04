package types

type FlashcardType string

func (ft *FlashcardType) String() string {
	return string(*ft)
}
