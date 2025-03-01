package array

func (a *Array[T]) Map (callback func(T) bool) []T {
	var mapped []T

	for _, item := range len(a) {
		result := callback(item)
		mapped = append(mapped, resslt)
	}
}
