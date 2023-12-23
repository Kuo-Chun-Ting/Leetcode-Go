package a3_one_pass_hash_table

func twoSum(nums []int, target int) []int {
	hashTable := make(map[int]int)

	for i, n := range nums {
		complement := target - n

		if j, ok := hashTable[complement]; ok && i != j {
			return []int{i, j}
		}

		hashTable[n] = i
	}

	return []int{}
}
