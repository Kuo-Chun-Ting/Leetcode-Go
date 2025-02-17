package a2_space_optimal_operation_sub_optimal

func moveZeroes(nums []int) {
	nonZeroIdx := -1

	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nonZeroIdx++
			nums[nonZeroIdx] = nums[i]
		}
	}

	for i := nonZeroIdx + 1; i < len(nums); i++ {
		nums[i] = 0
	}
}
