package leetcode

import (
	"strings"
	"unicode"
)

func isPalindrome2(s string) bool {
	res := ""
	s = strings.ToLower(s)
	for i := 0; i < len(s); i++ {
		if unicode.IsLetter(rune(s[i])) || unicode.IsNumber(rune(s[i])) {
			res += string(s[i])
		}
	}

	if len(res) == 0 {
		return true
	}

	for i := 0; i < len(res)/2; i++ {
		if res[i] != res[len(res)-i-1] {
			return false
		}
	}

	return true
}
