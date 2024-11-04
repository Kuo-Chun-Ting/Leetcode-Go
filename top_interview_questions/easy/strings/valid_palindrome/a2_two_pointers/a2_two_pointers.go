package a2_two_pointers

import "strings"

func isPalindrome(s string) bool {
	alphanumeric := filterAlphanumeric(s)
	upperStr := strings.ToUpper(alphanumeric)

	startPtr := 0
	endPtr := len(upperStr) - 1

	for startPtr < endPtr {
		if upperStr[startPtr] != upperStr[endPtr] {
			return false
		}

		startPtr++
		endPtr--
	}

	return true
}

func filterAlphanumeric(s string) string {
	alphanumeric := []rune{}

	for _, char := range s {
		if 'a' <= char && char <= 'z' || 'A' <= char && char <= 'Z' || '0' <= char && char <= '9' {
			alphanumeric = append(alphanumeric, char)
		}
	}
	return string(alphanumeric)
}
