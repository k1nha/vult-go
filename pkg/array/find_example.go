package array

import "fmt"

type Car struct {
	Id   int
	Name string
}

func FindExample() {
	arr := Array[Car]{
		Items: []Car{
			Car{
				Id:   1,
				Name: "bmw",
			},
		},
	}
	find := arr.Find(func(c Car) bool {
		return c.Name == "bmw"
	})
	fmt.Print(find.Name) // bmw
}
