package array

import "fmt"

func ReduceExample() {
	arr := ReduceArray[int, int]{
		Items: []int{1, 2, 3, 4, 5},
	}
	sum := arr.Reduce(func(accumulator, item int) int {
		return accumulator + item
	}, 0)
	fmt.Println(sum) // 15
}
