package euler

// Returns true if a number is the same forwards and backwards.
func IsPalindrome(num int) bool {
	if num < 0 {
		num *= -1
	}

	rev := 0
	tmp := num

	for tmp > 0 {
		dig := tmp % 10
		rev = rev*10 + dig
		tmp /= 10
	}

	return rev == num
}
