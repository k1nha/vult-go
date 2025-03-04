package array_test

import (
	"testing"

	"github.com/k1nha/vult-go/pkg/array"
)

type Product struct {
	Name  string
	Price int
}

func TestSumAllValues(t *testing.T) {
	arr := array.NewReduceArray[int, int](1, 2, 3, 4, 5)
	want := 15
	got := arr.Reduce(func(accumulator, item int) int {
		return accumulator + item
	}, 0)
	assertValues(t, want, got)
}

func TestSumPrices(t *testing.T) {
	products := []Product{
		{
			Name:  "Pencil",
			Price: 2,
		},
		{
			Name:  "Pen",
			Price: 4,
		},
	}
	arr := array.NewReduceArray[Product, int](products...)
	got := arr.Reduce(func(accumulator int, item Product) int {
		return item.Price + accumulator
	}, 0)
	want := 6
	assertValues(t, got, want)
}

func assertValues(t testing.TB, got, want int) {
	t.Helper()
	if want != got {
		t.Errorf("want %v got %v", want, got)
	}
}
