package array

func (a *Array[T]) Map(callback func(T) T) []T {
	var mapped []T

	for _, item := range a.Items {
		result := callback(item)
		mapped = append(mapped, result)
	}
	return mapped
}
