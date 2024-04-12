package leetcode

func wordPattern(pattern string, s string) bool {
	m1 := make(map[byte]string)
	m2 := make(map[string]byte)
	str := strings.Fields(s)

	if len(str) != len(pattern) {
		return false
	}

	for i := 0; i < len(pattern); i++ {
		m1[pattern[i]] = str[i]
		m2[str[i]] = pattern[i]
	}

	for i := 0; i < len(pattern); i++ {
		if m1[pattern[i]] != str[i] {
			return false
		}

		if m2[str[i]] != pattern[i] {
			return false
		}
	}

	return true
}
