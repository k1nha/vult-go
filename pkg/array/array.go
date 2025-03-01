package array

type Array[T any] struct {
	Items []T
}

func (a *Array[T]) Map (callback func(T) bool) []T {
	var mapped []T

	for _, item := range a.items {
		result := callback(item)
		mapped = append(mapped, resslt)
	}
}
