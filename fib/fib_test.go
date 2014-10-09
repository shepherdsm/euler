package fib

import (
	"reflect"
	"testing"
)

type testNth struct {
	n    uint
	want uint
}

func TestNth(t *testing.T) {
	var tests = []testNth{
		{0, 0},
		{1, 1},
		{2, 1},
		{10, 55},
		{50, 12586269025},
	}

	for _, pair := range tests {
		got := Nth(pair.n)

		if got != pair.want {
			t.Error(
				"For", pair.n,
				"got", got,
				"expected", pair.want,
			)
		}
	}
}

type testUnder struct {
	upper uint
	want  []uint
}

func TestUpper(t *testing.T) {
	var tests = []testUnder{
		{0, nil},
		{1, []uint{0}},
		{10, []uint{0, 1, 1, 2, 3, 5, 8}},
	}

	for _, pair := range tests {
		got := Under(pair.upper)

		if !reflect.DeepEqual(got, pair.want) {
			t.Error(
				"For", pair.upper,
				"got", got,
				"expected", pair.want,
			)
		}
	}
}
