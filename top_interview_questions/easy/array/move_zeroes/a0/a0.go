package a0

func moveZeroes(nums []int) {
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			for j := i + 1; j < len(nums); j++ {
				if nums[j] != 0 {
					temp := nums[i]
					nums[i] = nums[j]
					nums[j] = temp
					break
				}
			}
		}
	}
}
