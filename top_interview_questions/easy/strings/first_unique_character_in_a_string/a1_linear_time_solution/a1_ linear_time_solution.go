package a1_linear_time_solution

func firstUniqChar(s string) int {
	strCounter := make(map[rune]int)

	for _, char := range s {
		if _, ok := strCounter[char]; ok {
			strCounter[char]++
		} else {
			strCounter[char] = 1
		}
	}

	for i, char := range s {
		if strCounter[char] == 1 {
			return i
		}
	}

	return -1
}
