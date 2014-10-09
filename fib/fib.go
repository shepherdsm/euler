// Package Fib provides tools for working with fibonacci numbers.
package fib

// Nth returns the Nth fibonacci number.
func Nth(n uint) (uint) {
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
