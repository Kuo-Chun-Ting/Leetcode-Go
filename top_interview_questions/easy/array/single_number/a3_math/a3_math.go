package a3_math

func singleNumber(nums []int) int {
	totalSum := 0
	hashMap := make(map[int]bool, len(nums))
	for i := 0; i < len(nums); i++ {
		hashMap[nums[i]] = true
		totalSum += nums[i]
	}

	uniqueNumSum := 0
	for k, _ := range hashMap {
		uniqueNumSum += k
	}

	return 2*uniqueNumSum - totalSum
}
