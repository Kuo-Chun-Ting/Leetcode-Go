package remove_duplicates_from_sorted_array

import "fmt"

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	uniqueNumIndex := 0
	for j := uniqueNumIndex + 1; j < len(nums); j++ {
		if nums[uniqueNumIndex] != nums[j] {
			uniqueNumIndex++
			nums[uniqueNumIndex] = nums[j]
		}
	}

	return uniqueNumIndex + 1
}

func Run() {
	sortedArray := []int{1, 1, 2, 2, 3}
	result := removeDuplicates(sortedArray)
	fmt.Println(result)
}
