package a0

func romanToInt(s string) int {
	romanMap := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	sum := 0
	i := 0
	for i < len(s) {
		if i == len(s)-1 || romanMap[s[i]] >= romanMap[s[i+1]] {
			sum += romanMap[s[i]]
			i += 1
		} else {
			sum += romanMap[s[i+1]] - romanMap[s[i]]
			i += 2
		}
	}

	return sum
}
