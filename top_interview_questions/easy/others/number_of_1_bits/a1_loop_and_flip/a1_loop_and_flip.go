package a1_loop_and_flip

func hammingWeight(n int) int {
	bits := 0
	mask := 1

	for i := 0; i < 32; i++ {
		andResult := n & mask
		if andResult != 0 {
			bits++
		}
		mask = mask << 1
	}

	return bits
}
