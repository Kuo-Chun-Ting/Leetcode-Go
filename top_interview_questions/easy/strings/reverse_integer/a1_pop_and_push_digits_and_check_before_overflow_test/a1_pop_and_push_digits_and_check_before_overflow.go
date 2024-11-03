package a1_pop_and_push_digits_and_check_before_overflow

import (
	"math"
)

func reverse(x int) int {
	var n float64 = 31
	maxInt := math.Pow(2, n) - 1
	minInt := math.Pow(-2, n)
	rev := 0

	for x != 0 {
		pop := x % 10
		x = x / 10

		if rev > int(maxInt)/10 || (rev == int(maxInt)/10 && pop > int(maxInt)%10) {
			return 0
		}

		if rev < int(minInt)/10 || (rev == int(minInt)/10 && pop < int(minInt)%10) {
			return 0
		}

		rev = rev*10 + pop
	}

	return rev
}
