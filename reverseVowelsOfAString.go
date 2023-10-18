package leetcode

func reverseVowels(s string) string {
	vowels := ""
	res := ""
	for i := 0; i < len(s); i++ {
		if isVowel(s[i]) {
			vowels += string(s[i])
		}
	}

	index := len(vowels) - 1
	for i := 0; i < len(s); i++ {
		if isVowel(s[i]) {
			res += string(vowels[index])
			index--
		} else {
			res += string(s[i])
		}
	}

	return res
}

func isVowel(s byte) bool {
	if s == 'a' || s == 'e' || s == 'i' || s == 'o' || s == 'u' || s == 'A' || s == 'E' || s == 'I' || s == 'O' || s == 'U' {
		return true
	}
	return false
}
