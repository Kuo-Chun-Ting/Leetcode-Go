package a0_mine

func plusOne(digits []int) []int {
	i := len(digits) - 1

	for {
		digits[i] += 1

		if digits[i] < 10 {
			break
		}

		digits[i] = digits[i] % 10
		i--

		if i < 0 {
			digits = append([]int{1}, digits...)
			break
		}
	}

	return digits
}
