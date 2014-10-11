package fib

import (
	"reflect"
	"testing"
)

var funcVals = []uint{
	0,
	1,
	2,
	4,
	5,
	7,
	10,
	16,
	20,
	44,
}

var fibNth = []uint{
	0,
	1,
	1,
	3,
	5,
	13,
	55,
	987,
	6765,
	701408733,
}

var fibUnder = [][]uint{
	nil,
	nil,
	{1, 1},
	{1, 1, 2, 3},
	{1, 1, 2, 3},
	{1, 1, 2, 3, 5},
	{1, 1, 2, 3, 5, 8},
	{1, 1, 2, 3, 5, 8, 13},
	{1, 1, 2, 3, 5, 8, 13},
	{1, 1, 2, 3, 5, 8, 13, 21, 34},
}

func TestNth(t *testing.T) {
	for i, num := range funcVals {
		if got := Nth(num); got != fibNth[i] {
			t.Errorf("Nth(%d) = %d, want %d", num, got, fibNth[i])
		}
	}
}

func TestUnder(t *testing.T) {
	for i, num := range funcVals {
		if got := Under(num); !reflect.DeepEqual(got, fibUnder[i]) {
			t.Errorf("Under(%d) = %v, want %v", num, got, fibUnder[i])
		}
	}
}
