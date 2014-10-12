package euler

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

var collatz = [][]uint{
	nil,
	{1},
	{2, 1},
	{4, 2, 1},
	{5, 16, 8, 4, 2, 1},
	{7, 22, 11, 34, 17, 52, 26, 13, 40, 20, 10, 5, 16, 8, 4, 2, 1},
	{10, 5, 16, 8, 4, 2, 1},
	{16, 8, 4, 2, 1},
	{20, 10, 5, 16, 8, 4, 2, 1},
	{44, 22, 11, 34, 17, 52, 26, 13, 40, 20, 10, 5, 16, 8, 4, 2, 1},
}

var nextnum = []uint{
	0,
	4,
	1,
	2,
	16,
	22,
	5,
	8,
	10,
	22,
}

func TestCollatz(t *testing.T) {
	for i, num := range funcVals {
		if got := Collatz(num); !reflect.DeepEqual(got, collatz[i]) {
			t.Errorf("Collatz(%d) = %v, want %v", num, got, collatz[i])
		}
	}
}

func TestNextNum(t *testing.T) {
	for i, num := range funcVals {
		if got := nextNum(num); got != nextnum[i] {
			t.Errorf("nextNum(%d) = %d, want %d", num, got, nextnum[i])
		}
	}
}
