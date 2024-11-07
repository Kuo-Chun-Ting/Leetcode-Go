package a1_follow_the_rules

import "math"

func myAtoi(s string) int {
	sign := 1
	result := 0
	index := 0
	n := len(s)

	for index < n && s[index] == ' ' {
		index++
	}

	if index < n && s[index] == '+' {
		sign = 1
		index++
	} else if index < n && s[index] == '-' {
		sign = -1
		index++
	}

	for index < n && s[index] >= '0' && s[index] <= '9' {
		digit := int(s[index] - '0')

		if result > math.MaxInt32/10 || (result == math.MaxInt32/10 && digit > math.MaxInt32%10) {
			if sign == 1 {
				return math.MaxInt32
			}
			return math.MinInt32
		}

		result = 10*result + digit
		index++
	}

	return sign * result
}
