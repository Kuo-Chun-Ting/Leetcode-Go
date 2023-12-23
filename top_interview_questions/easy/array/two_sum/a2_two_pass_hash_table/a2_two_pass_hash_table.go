package a2_two_pass_hash_table

func twoSum(nums []int, target int) []int {
	hashTable := make(map[int]int)

	for i, n := range nums {
		hashTable[n] = i
	}

	for i, n := range nums {
		complement := target - n
		if j, ok := hashTable[complement]; ok && j != i {
			return []int{i, j}
		}
	}

	return []int{}
}
