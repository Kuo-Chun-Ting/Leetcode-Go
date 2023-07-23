package reverse_string

import "fmt"

func reverseString(s []byte) {
	if len(s) == 0 {
		return
	}

	i := len(s)
	for j := 0; j < i; j++ {
		if i <= j {
			break
		}

		temp := s[j]
		s[j] = s[i-1]
		s[i-1] = temp
		i--
	}
}

func Run() {
	str := []byte{'h', 'e', 'l', 'l', 'o'}
	reverseString(str)
	fmt.Println(str)
}
