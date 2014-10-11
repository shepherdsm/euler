package euler

func Collatz(seed uint) []uint {
	if seed <= 0 {
		return nil
	}

	var results = []uint{seed}

	for seed != 1 {
		seed = nextNum(seed)
		results = append(results, seed)
	}

	return results
}

func nextNum(n uint) uint {
	if n%2 == 0 {
		return n / 2
	} else {
		return 3*n + 1
	}
}
