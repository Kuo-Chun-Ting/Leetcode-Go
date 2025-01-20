package a2_three_pointers_start_from_the_beginning

func merge(nums1 []int, m int, nums2 []int, n int) {
	copiedNums1 := make([]int, len(nums1))
	copy(copiedNums1, nums1)

	currIdx, ptr1, ptr2 := 0, 0, 0

	for ptr1 < m && ptr2 < n {
		if copiedNums1[ptr1] < nums2[ptr2] {
			nums1[currIdx] = copiedNums1[ptr1]
			ptr1++
		} else {
			nums1[currIdx] = nums2[ptr2]
			ptr2++
		}

		currIdx++
	}

	for i := currIdx; i < len(nums1); i++ {
		if ptr1 == m {
			nums1[i] = nums2[ptr2]
			ptr2++
		} else if ptr2 == n {
			nums1[i] = copiedNums1[ptr1]
			ptr1++
		}
	}
}
