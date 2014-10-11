// Package primes handles anything to do with prime numbers from generating to testing.
package primes

import "math"

// IsPrime determines if n is a prime number.
func IsPrime(n uint) bool {
	if n == 2 {
		return true
	}
	if n <= 1 || n%2 == 0 {
		return false
	}

	upper := uint(math.Sqrt(float64(n)))

	for k := uint(3); k <= upper; k += 2 {
		if n%k == 0 {
			return false
		}
	}

	return true
}

// Under generates all primes in the range [2,n).
func Under(n uint) []uint {
	if n <= 2 {
		return nil
	}

	primes := []uint{2}

	for k := uint(3); k < n; k += 2 {
		if IsPrime(k) {
			primes = append(primes, k)
		}
	}

	return primes
}

// sieveEratosthenes uses the Sieve of Eratosthenes to generate primes
// in the range [2,n).
func sieveEratosthenes(n uint) []uint {
	var primes []uint
	sieve := make(map[uint]bool)

	for i := uint(2); i < n; i++ {
		sieve[i] = true
	}

	for i := uint(2); i <= n; i++ {
		if cand := sieve[i]; cand {
			primes = append(primes, i)
			for k := 2 * i; k <= n; k += i {
				sieve[k] = false
			}
		}
	}

	return primes
}
