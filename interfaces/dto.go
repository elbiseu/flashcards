package interfaces

type DTO interface {
	ToMap() map[string]any
}
