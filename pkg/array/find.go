package array

func (a *Array[T]) Find(callback func(T) bool) T {
	var nil T
	for _, item := range a.Items {
		if callback(item) {
			return item
		}
	}
	return nil
}
