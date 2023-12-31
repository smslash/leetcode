package leetcode

func maxLengthBetweenEqualCharacters(s string) int {
	max := -1
	for i := 0; i < len(s); i++ {
		for j := len(s) - 1; j > i; j-- {
			if s[i] == s[j] && max < len(s[i+1:j]) {
				max = len(s[i+1 : j])
			}
		}
	}

	return max
}
