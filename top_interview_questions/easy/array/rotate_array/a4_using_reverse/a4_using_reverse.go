package a4_using_reverse

// Check twice
func rotate(nums []int, k int) {
	n := len(nums)
	k = k % n
	reverse(nums)
	reverse(nums[0:k])
	reverse(nums[k:n])
}

func reverse(nums []int) {
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		temp := nums[i]
		nums[i] = nums[j]
		nums[j] = temp
	}
}
