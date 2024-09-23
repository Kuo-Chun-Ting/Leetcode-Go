package a2_hash_set

func missingNumber(nums []int) int {
	set := make(map[int]struct{})
	for _, n := range nums {
		set[n] = struct{}{}
	}

	for i := 0; i < len(nums)+1; i++ {
		if _, ok := set[i]; !ok {
			return i
		}
	}

	return 0
}
