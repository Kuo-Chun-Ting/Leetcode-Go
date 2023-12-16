package a3_using_cyclic_replacements

func rotate(nums []int, k int) {
	n := len(nums)
	start, idxAt := 0, 0
	temp := -1
	curr := nums[0]

	for i := 0; i < n; i++ {
		jumpTo := (idxAt + k) % n
		temp = nums[jumpTo]
		nums[jumpTo] = curr

		if jumpTo == start {
			start, idxAt = (jumpTo+1)%n, (jumpTo+1)%n
			curr = nums[idxAt]
		} else {
			idxAt = jumpTo
			curr = temp
		}
	}
}
