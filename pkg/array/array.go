package array

type Array[T any] struct {
	Items []T
}

func NewArray[T any](items ...T) *Array[T] {
	return &Array[T]{
		Items: items,
	}
}
