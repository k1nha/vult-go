package array

func (a *Array[T]) Filter(callback func(T) bool) []T {
	var filtered []T
	for _, item := range a.Items {
		if callback(item) {
			filtered = append(filtered, item)
		}
	}
	return filtered
}
