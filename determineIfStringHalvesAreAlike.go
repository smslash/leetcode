package leetcode

func halvesAreAlike(s string) bool {
	n := len(s)
	l, r := 0, 0

	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		l += isVowel(s[i])
		r += isVowel(s[j])
	}
	return l == r
}

func isVowel(x byte) int {
	if 'a' == x || 'e' == x || 'i' == x || 'o' == x || 'u' == x {
		return 1
	}

	if 'A' == x || 'E' == x || 'I' == x || 'O' == x || 'U' == x {
		return 1
	}
    
	return 0
}
