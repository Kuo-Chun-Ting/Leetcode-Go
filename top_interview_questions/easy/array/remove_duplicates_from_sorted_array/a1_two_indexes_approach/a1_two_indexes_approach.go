package a1_two_indexes_approach

// Check twice
func removeDuplicates(nums []int) int {
	insertedIdx := 1

	for i := insertedIdx; i < len(nums); i++ {
		if nums[i-1] != nums[i] {
			nums[insertedIdx] = nums[i]
			insertedIdx++
		}
	}

	return insertedIdx
}
