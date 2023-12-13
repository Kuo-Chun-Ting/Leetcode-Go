package brute_force

func rotate(nums []int, k int) {
	for i := 0; i > k; i++ {
		prv := nums[len(nums)-1]
		for j := 0; j < len(nums); j++ {
			temp := nums[j]
			nums[j] = prv
			prv = temp
		}
	}
}
