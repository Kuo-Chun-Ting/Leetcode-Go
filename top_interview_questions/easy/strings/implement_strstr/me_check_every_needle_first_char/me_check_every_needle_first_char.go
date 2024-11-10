package me_check_every_needle_first_char

func strStr(haystack string, needle string) int {
	firstNeedleCharIdxList := findAllFirstNeedleCharInHaystack(haystack, needle)

	for _, idx := range firstNeedleCharIdxList {
		hayIdx := idx
		neeIdx := 0

		for neeIdx < len(needle) {
			if haystack[hayIdx] == needle[neeIdx] {
				hayIdx++
				neeIdx++
			} else {
				break
			}

			if neeIdx == len(needle) {
				return idx
			}
		}
	}

	return -1
}

func findAllFirstNeedleCharInHaystack(haystack string, needle string) []int {
	firstNeedleCharIdxList := []int{}

	for i, char := range haystack {
		if char == rune(needle[0]) && len(haystack[i:]) >= len(needle) {
			firstNeedleCharIdxList = append(firstNeedleCharIdxList, i)
		}
	}

	return firstNeedleCharIdxList
}
