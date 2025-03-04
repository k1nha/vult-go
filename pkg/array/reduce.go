package array

type ReduceArray[T any, R any] struct {
	Items []T
}

func NewReduceArray[T any, R any](items ...T) *ReduceArray[T, R] {
	return &ReduceArray[T, R]{
		Items: items,
	}
}

func (a *ReduceArray[T, R]) Reduce(callback func(accumulator R, item T) R, initial R) R {
	accumulator := initial
	for _, item := range a.Items {
		accumulator = callback(accumulator, item)
	}
	return accumulator
}
