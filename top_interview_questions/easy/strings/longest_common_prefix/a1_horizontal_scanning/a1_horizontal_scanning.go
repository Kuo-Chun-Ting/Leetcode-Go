package a1_horizontal_scanning

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	longestCommonPrefix := strs[0]
	firstStr := strs[0]

	for _, s := range strs[1:] {
		commonPrefix := findCommonPrefix(firstStr, s)
		if len(commonPrefix) < len(longestCommonPrefix) {
			longestCommonPrefix = commonPrefix
		}
	}
	return longestCommonPrefix
}

func findCommonPrefix(str1 string, str2 string) string {
	index := 0
	for index < len(str1) && index < len(str2) {
		if str1[index] == str2[index] {
			index++
			continue
		} else {
			break
		}
	}

	return str1[:index]
}
