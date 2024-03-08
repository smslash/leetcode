package leetcode

func balancedStringSplit(s string) int {
	m := make(map[byte]int)
	count := 0
	for i := 0; i < len(s); i++ {
		m[s[i]]++

		if m['R'] == m['L'] {
			m['R'] = 0
			m['L'] = 0
			count++
		}
	}
	return count
}
