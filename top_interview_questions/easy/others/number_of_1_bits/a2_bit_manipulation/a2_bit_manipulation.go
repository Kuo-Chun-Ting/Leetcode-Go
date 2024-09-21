package a2_bit_manipulation

func hammingWeight(n int) int {
	bits := 0

	for {
		if n == 0 {
			break
		}

		n = n & (n - 1)
		bits++
	}

	return bits
}
