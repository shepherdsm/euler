package euler

import ("testing"
"reflect"
)

type testCollatz struct {
    seed uint
    want []uint
}

func TestCollatz(t *testing.T) {
    var tests = []testCollatz {
        {1, []uint {1}},
        {10, []uint {10, 5, 16, 8, 4, 2, 1}},
    }

    for _, pair := range tests {
        got := Collatz(pair.seed)

        if !reflect.DeepEqual(got, pair.want) {
            t.Error(
                "For", pair.seed,
                "got", got,
                "expected", pair.want,
            )
        }
    }
}
