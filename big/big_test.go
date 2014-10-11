package big

import (
	"math/big"
	"testing"
)

func pow(base, exp int64) *big.Int {
	return big.NewInt(0).Exp(big.NewInt(base), big.NewInt(exp), nil)
}

var digitSumVals = []*big.Int{
	pow(2, 10),
	pow(2, 100),
	pow(2, 1000),
	pow(2, 10000),
	pow(2, 100000),
}

// Answers computed using wolframalpha.com
var digitsum = []int64{
	7,
	115,
	1366,
	13561,
	135178,
}

func TestDigitSum(t *testing.T) {
	for i, num := range digitSumVals {
		if got := DigitSum(num); got != digitsum[i] {
			t.Errorf("DigitSum(%v) = %d, want %d", num, got, digitsum[i])
		}
	}
}

func benchDigitSum(i *big.Int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		DigitSum(i)
	}
}

func BenchmarkDigitSum2_10(b *testing.B)     { benchDigitSum(digitSumVals[0], b) }
func BenchmarkDigitSum2_100(b *testing.B)    { benchDigitSum(digitSumVals[1], b) }
func BenchmarkDigitSum2_1000(b *testing.B)   { benchDigitSum(digitSumVals[2], b) }
func BenchmarkDigitSum2_10000(b *testing.B)  { benchDigitSum(digitSumVals[3], b) }
func BenchmarkDigitSum2_100000(b *testing.B) { benchDigitSum(digitSumVals[4], b) }
