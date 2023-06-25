package remove_duplicates_from_sorted_array

import "fmt"

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	i := 0
	for j := i + 1; j < len(nums); j++ {
		if nums[i] != nums[j] {
			i++
			nums[i] = nums[j]
		}
	}

	return i + 1
}

func Run() {
	sortedArray := []int{1, 1, 2, 2, 3}
	result := removeDuplicates(sortedArray)
	fmt.Println(result)
}
