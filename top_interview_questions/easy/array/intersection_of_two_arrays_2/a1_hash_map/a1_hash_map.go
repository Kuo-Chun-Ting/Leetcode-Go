package a1_hash_map

// Check twice
func intersect(nums1 []int, nums2 []int) []int {
	nums1Map := make(map[int]int)

	for _, n := range nums1 {
		nums1Map[n] += 1
	}

	result := []int{}
	for _, n := range nums2 {
		if v, ok := nums1Map[n]; ok && v != 0 {
			result = append(result, n)
			nums1Map[n] -= 1
		}
	}
	return result
}
