package a1_brute_force

func climbStairs(n int) int {
	return climbStairsFrom(0, n)
}

func climbStairsFrom(i, n int) int {
	if i > n {
		return 0
	}

	if i == n {
		return 1
	}

	return climbStairsFrom(i+1, n) + climbStairsFrom(i+2, n)
}
