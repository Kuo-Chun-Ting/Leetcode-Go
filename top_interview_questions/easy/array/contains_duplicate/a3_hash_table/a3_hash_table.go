package a3_hash_table

// Check twice
func containsDuplicate(nums []int) bool {
	hashMap := map[int]bool{}

	for i := 0; i < len(nums); i++ {
		if _, ok := hashMap[nums[i]]; ok {
			return true
		}

		hashMap[nums[i]] = true
	}

	return false
}
