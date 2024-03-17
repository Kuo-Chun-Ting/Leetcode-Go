package a2_dynamic_programming_kadanes_algorithm

func maxSubArray(nums []int) int {
	maxSubArrSum, currSubArrSum := nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] > currSubArrSum+nums[i] {
			currSubArrSum = nums[i]
		} else {
			currSubArrSum += nums[i]
		}

		if currSubArrSum > maxSubArrSum {
			maxSubArrSum = currSubArrSum
		}
	}

	return maxSubArrSum
}
