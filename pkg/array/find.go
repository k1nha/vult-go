package array

func (a *Array[T]) Find(callback func(T) bool) T {
	var zero T
	for _, item := range a.Items {
		if callback(item) {
			return item
		}
	}
	return zero
}
