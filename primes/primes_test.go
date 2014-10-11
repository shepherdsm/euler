package primes

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
	47,
}

var prime = []bool{
	false,
	false,
	true,
	false,
	true,
	true,
	false,
	false,
	false,
	true,
}

var primes = [][]uint{
	nil,
	nil,
	nil,
	{2, 3},
	{2, 3},
	{2, 3, 5},
	{2, 3, 5, 7},
	{2, 3, 5, 7, 11, 13},
	{2, 3, 5, 7, 11, 13, 17, 19},
	{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43},
}

func TestIsPrime(t *testing.T) {
	for i, num := range funcVals {
		if got := IsPrime(num); got != prime[i] {
			t.Errorf("isPrime(%d) = %t, want %t", num, got, prime[i])
		}
	}
}

func TestSieveEratosthenes(t *testing.T) {
	for i, num := range funcVals {
		if got := sieveEratosthenes(num); !reflect.DeepEqual(got, primes[i]) {
			t.Errorf("sieveEratosthenes(%d) = %v, want %v", num, got, primes[i])
		}
	}
}

func TestUnder(t *testing.T) {
	for i, num := range funcVals {
		if got := Under(num); !reflect.DeepEqual(got, primes[i]) {
			t.Errorf("Under(%d) = %v, want %v", num, got, primes[i])
		}
	}
}

func benchSieveEratosthenes(i uint, b *testing.B) {
	for n := 0; n < b.N; n++ {
		sieveEratosthenes(i)
	}
}

func BenchmarkSieveEratosthenes10(b *testing.B)      { benchSieveEratosthenes(10, b) }
func BenchmarkSieveEratosthenes100(b *testing.B)     { benchSieveEratosthenes(100, b) }
func BenchmarkSieveEratosthenes1000(b *testing.B)    { benchSieveEratosthenes(1000, b) }
func BenchmarkSieveEratosthenes10000(b *testing.B)   { benchSieveEratosthenes(10000, b) }
func BenchmarkSieveEratosthenes100000(b *testing.B)  { benchSieveEratosthenes(100000, b) }
func BenchmarkSieveEratosthenes1000000(b *testing.B) { benchSieveEratosthenes(1000000, b) }

func benchIsPrime(i uint, b *testing.B) {
	for n := 0; n < b.N; n++ {
		IsPrime(i)
	}
}

func BenchmarkIsPrime10(b *testing.B)       { benchIsPrime(29, b) }
func BenchmarkIsPrime100(b *testing.B)      { benchIsPrime(541, b) }
func BenchmarkIsPrime1000(b *testing.B)     { benchIsPrime(7919, b) }
func BenchmarkIsPrime10000(b *testing.B)    { benchIsPrime(104729, b) }
func BenchmarkIsPrime100000(b *testing.B)   { benchIsPrime(1299827, b) }
func BenchmarkIsPrime1000000(b *testing.B)  { benchIsPrime(15485863, b) }
func BenchmarkIsPrime10000000(b *testing.B) { benchIsPrime(179424673, b) }

func benchUnder(i uint, b *testing.B) {
	for n := 0; n < b.N; n++ {
		Under(i)
	}
}

func BenchmarkUnder10(b *testing.B)      { benchUnder(10, b) }
func BenchmarkUnder100(b *testing.B)     { benchUnder(100, b) }
func BenchmarkUnder1000(b *testing.B)    { benchUnder(1000, b) }
func BenchmarkUnder10000(b *testing.B)   { benchUnder(10000, b) }
func BenchmarkUnder100000(b *testing.B)  { benchUnder(100000, b) }
func BenchmarkUnder1000000(b *testing.B) { benchUnder(1000000, b) }
