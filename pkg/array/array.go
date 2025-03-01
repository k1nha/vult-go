package array

type Array[T any] struct {
	Items []T
}

func NewArr() Array[T] {
	return Array[T]{}
}

