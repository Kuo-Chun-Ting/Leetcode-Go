package a1_stacks

func isValid(s string) bool {
	parenthsisMapping := map[rune]rune{')': '(', ']': '[', '}': '{'}
	stack := []rune{}

	for _, currChar := range s {
		if correctParenthsis, ok := parenthsisMapping[currChar]; ok {
			if len(stack) == 0 {
				return false
			}

			var topChar rune
			topChar, stack = pop(stack)
			if topChar != correctParenthsis {
				return false
			}
		} else {
			stack = push(stack, currChar)
		}
	}

	return len(stack) == 0
}

func push(stack []rune, value rune) []rune {
	return append(stack, value)
}

func pop(stack []rune) (rune, []rune) {
	return stack[len(stack)-1], stack[:len(stack)-1]
}
