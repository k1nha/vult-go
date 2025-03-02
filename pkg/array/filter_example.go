package array

import "fmt"

func FilterExample() {
	arry := Array[int]{
		Items: []int{1, 2, 4},
	}
	filtered := arry.Filter(func(i int) bool {
		return i%2 == 0
	})
	for i := range len(filtered) {
		fmt.Println(filtered[i]) // 2,4
	}
}
