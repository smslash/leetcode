package leetcode

func firstPalindrome(words []string) string {
	for i := 0; i < len(words); i++ {
		if isPalindrome(words[i]) {
			return words[i]
		}
	}

	return ""
}

func isPalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}