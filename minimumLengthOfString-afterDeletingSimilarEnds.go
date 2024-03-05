package leetcode

func minimumLength(s string) int {
	if len(s) < 2 {
		return len(s)
	}

	count := 0
	i := 0
	j := len(s) - 1

	for i < j {
		if s[i] == s[j] {
			count += 2
			i++
			j--
			for i < j && s[i] == s[i-1] {
				count++
				i++
			}
			for i <= j && s[j] == s[j+1] {
				count++
				j--
			}
		} else {
			break
		}
	}

	return len(s) - count
}
