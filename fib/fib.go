// Package Fib provides tools for working with fibonacci numbers.
package fib

// Nth returns the Nth fibonacci number.
func Nth(n uint) uint {
	if n == 0 {
		return 0
	}

	var nPrev uint = 0
	var nCur uint = 1
	var i uint = 0

	for i = 1; i < n; i++ {
		tmp := nCur

		nCur += nPrev
		nPrev = tmp
	}

	return nCur
}

// Generates all the fibonacci numbers under, but not including, upper.
func Under(upper uint) []uint {
	var results []uint
	var n uint = 1

	fib := Nth(n)

	for fib < upper {
		results = append(results, fib)
		n++
		fib = Nth(n)
	}

	return results
}
