package interfaces

type Store interface {
	Put(entity Entity) error
	Get(entity Entity) error
	Remove(entity Entity) error
}
