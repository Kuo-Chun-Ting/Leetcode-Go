package a2_frequency_counter

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	alphabetCounter := make(map[rune]int, 26)

	for i := 0; i < len(s); i++ {
		sChar := rune(s[i])
		tChar := rune(t[i])

		alphabetCounter[sChar]++
		alphabetCounter[tChar]--
	}

	for _, count := range alphabetCounter {
		if count != 0 {
			return false
		}
	}

	return true
}
