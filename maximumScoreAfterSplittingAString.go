package leetcode

func maxScore(s string) int {
	max := 0
	for i := 1; i < len(s); i++ {
		count := 0
		
		for j := 0; j < len(s[:i]); j++ {
			if s[:i][j] == '0' {
				count++
			}
		}

		for j := 0; j < len(s[i:]); j++ {
			if s[i:][j] == '1' {
				count++
			}
		}

		if max < count {
			max = count
		}
	}

	return max
}
