package a2_bit_manipulation

func hammingWeight(n int) int {
	bits := 0
	check := n

	for {
		if check == 0 {
			break
		}

		check = check & (check - 1)
		bits++
	}

	return bits
}
