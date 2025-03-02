package array_test

import (
	"reflect"
	"testing"

	"github.com/k1nha/vult-go/pkg/array"
)

func TestMapWithStrings(t *testing.T) {
	stringArray := array.Array[string]{Items: []string{"a", "b", "c"}}
	mappedStrings := stringArray.Map(func(s string) string {
		return s + s
	})
	expectedStrings := []string{"aa", "bb", "cc"}
	if !reflect.DeepEqual(mappedStrings, expectedStrings) {
		t.Errorf("Expected %v, but got %v", expectedStrings, mappedStrings)
	}
}

func TestMap(t *testing.T) {
	got := array.Array[int]{
		Items: []int{1, 2, 3},
	}
	mapped := got.Map(func(i int) int {
		return i * i
	})
	expect := []int{1, 4, 9}
	if !reflect.DeepEqual(mapped, expect) {
		t.Errorf("Expected %v, but got %v", expect, mapped)
	}
}
