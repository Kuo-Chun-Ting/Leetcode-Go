package a1_loop_iteration

func isPowerOfThree(n int) bool {
	if n < 1 {
		return false
	}

	for n%3 == 0 {
		n = n / 3
	}

	return n == 1
}
