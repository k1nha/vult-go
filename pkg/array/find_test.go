package array_test

import (
	"reflect"
	"testing"

	"github.com/k1nha/vult-go/pkg/array"
)

type Car struct {
	Id   int
	Name string
}

var cars = []Car{
	{
		Id:   1,
		Name: "bmw",
	},
	{
		Id:   2,
		Name: "audi",
	},
}

func TestFind(t *testing.T) {
	arr := array.NewArray(cars...)
	find := arr.Find(func(c Car) bool {
		return c.Name == "bmw"
	})
	expected := arr.Items[0]
	if !reflect.DeepEqual(find, expected) {
		t.Errorf("Expected %v, but got %v", find, expected)
	}
}

func TestCannotFind(t *testing.T) {
	arr := array.NewArray(cars...)
	find := arr.Find(func(c Car) bool {
		return c.Name == "mercedes"
	})
	expected := arr.Items[0]

	if reflect.DeepEqual(find, expected) {
		t.Errorf("Expected %v, but got %v", find, expected)
	}
}
