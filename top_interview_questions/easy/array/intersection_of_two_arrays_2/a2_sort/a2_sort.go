package a2_sort

import "sort"

func intersect(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)

	i, j := 0, 0
	intersection := []int{}

	for {
		if i >= len(nums1) || j >= len(nums2) {
			break
		}

		if nums1[i] < nums2[j] {
			i++
		} else if nums1[i] > nums2[j] {
			j++
		} else {
			intersection = append(intersection, nums1[i])
			i++
			j++
		}

	}

	return intersection
}
