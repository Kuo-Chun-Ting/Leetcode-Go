package a1_optimized_brute_force

import "math"

func maxSubArray(nums []int) int {
	maxSum := math.MinInt

	for i := 0; i < len(nums); i++ {
		currSubArrSum := 0
		for j := i; j < len(nums); j++ {
			currSubArrSum += nums[j]
			if currSubArrSum > maxSum {
				maxSum = currSubArrSum
			}
		}
	}

	return maxSum
}
