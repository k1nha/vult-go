package array_test

import (
	"reflect"
	"testing"

	"github.com/k1nha/vult-go/pkg/array"
)

func TestFilter(t *testing.T) {
	arry := array.NewArray(1, 2, 3)
	filtered := arry.Filter(func(i int) bool {
		return i%2 == 0
	})
	expected := []int{2, 4}
	if !reflect.DeepEqual(filtered, expected) {
		t.Errorf("Expected %v, but got %v", expected, filtered)
	}
}

func TestFilterWithStrings(t *testing.T) {
	arry := array.Array[string]{
		Items: []string{"The Batman", "Joker", "Pinguim"},
	}
	filtered := arry.Filter(func(i string) bool {
		return i != "The Batman"
	})
	expected := []string{"Joker", "Pinguim"}
	if !reflect.DeepEqual(filtered, expected) {
		t.Errorf("Expected %v, but got %v", expected, filtered)
	}
}
