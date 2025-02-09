package a2_hash_table

func singleNumber(nums []int) int {
	hashMap := make(map[int]int, len(nums))

	for i := 0; i < len(nums); i++ {
		if _, ok := hashMap[nums[i]]; ok {
			hashMap[nums[i]]++
		} else {
			hashMap[nums[i]] = 1
		}
	}

	for k, v := range hashMap {
		if v == 1 {
			return k
		}
	}

	return 0
}
