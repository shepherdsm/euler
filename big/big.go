package big

import "math/big"

func DigitSum(num *big.Int) int64 {
	sum := big.NewInt(0)
	rem := big.NewInt(0)

	for num.Cmp(big.NewInt(0)) > 0 {
		num.DivMod(num, big.NewInt(10), rem)
		sum.Add(sum, rem)
	}

	return sum.Int64()
}
