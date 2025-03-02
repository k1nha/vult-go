package array

import "fmt"

func Example() {
	arry := Array[int]{
		Items: []int{1, 2, 4},
	}
	arryMultiple := arry.Map(func(i int) int {
		return i * 3
	})
	for i := range len(arryMultiple) {
		fmt.Println(arryMultiple[i]) // 2, 4, 8
	}
}
