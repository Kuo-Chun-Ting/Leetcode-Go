package a2_using_extra_array

func rotate(nums []int, k int) {
	extraArr := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		position := (i + k) % len(nums)
		extraArr[position] = nums[i]
	}

	copy(nums, extraArr)
}
