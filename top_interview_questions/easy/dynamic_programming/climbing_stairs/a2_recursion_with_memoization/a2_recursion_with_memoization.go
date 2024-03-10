package a2_recursion_with_memoization

func climbStairs(n int) int {
	memo := make([]int, n+1)
	return climbStairsFrom(0, n, &memo)
}

func climbStairsFrom(i, n int, memo *[]int) int {
	if i > n {
		return 0
	}

	if i == n {
		(*memo)[i] = 1
		return 1
	}

	if (*memo)[i] > 0 {
		return (*memo)[i]
	}

	(*memo)[i] = climbStairsFrom(i+1, n, memo) + climbStairsFrom(i+2, n, memo)
	return (*memo)[i]
}
