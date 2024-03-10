package a4_fibonacci_number

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}

	if n == 2 {
		return 2
	}

	first := 1
	second := 2
	for i := 3; i < n+1; i++ {
		third := first + second
		first = second
		second = third
	}
	return second
}
