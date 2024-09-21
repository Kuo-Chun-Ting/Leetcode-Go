package a1_bit_by_bit

func reverseBits(num uint32) uint32 {
	result := uint32(0)
	power := 31

	for {
		result += (num & 1) << power

		power--
		num = num >> 1

		if power < 0 {
			break
		}
	}
	return result
}
