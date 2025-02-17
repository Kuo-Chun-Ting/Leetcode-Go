package a1_schoolbook_addition_with_carry

// Check twice
func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i]+1 <= 9 {
			digits[i] += 1
			return digits
		}
		digits[i] = 0
	}

	return append([]int{1}, digits...)
}
