package a3_brian_kernighans_algorithm_test

func hammingDistance(x int, y int) int {
	diff := x ^ y
	bits := 0

	for {
		if diff == 0 {
			break
		}

		diff = diff & (diff - 1)
		bits++
	}

	return bits
}
