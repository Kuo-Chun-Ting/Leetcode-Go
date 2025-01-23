package a1_recursion_with_memoization

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	n := len(nums)
	maxRobbedAmount := make([]int, n)
	maxRobbedAmount[n-2], maxRobbedAmount[n-1] = max(nums[n-1], nums[n-2]), nums[n-1]

	for i := n - 3; i >= 0; i-- {
		maxRobbedAmount[i] = max(nums[i]+maxRobbedAmount[i+2], maxRobbedAmount[i+1])
	}

	return maxRobbedAmount[0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
