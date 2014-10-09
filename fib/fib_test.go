package fib

import "testing"

type testPair struct {
    n uint
    want uint
}

var tests = []testPair {
    {0, 0},
    {1, 1},
    {2, 1},
    {10, 55},
    {50, 12586269025},
}

func TestNth(t *testing.T) {
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
