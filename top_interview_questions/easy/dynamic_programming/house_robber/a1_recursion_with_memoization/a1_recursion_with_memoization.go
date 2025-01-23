package a1_recursion_with_memoization

func rob(nums []int) int {
	memo := make(map[int]int, len(nums))

	return robFrom(0, nums, memo)
}

func robFrom(i int, nums []int, memo map[int]int) int {
	if i >= len(nums) {
		return 0
	}

	if _, ok := memo[i]; ok {
		return memo[i]
	}

	sum := max(robFrom(i+1, nums, memo), nums[i]+robFrom(i+2, nums, memo))
	memo[i] = sum
	return sum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
