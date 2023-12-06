package leetcode

func countConsistentStrings(allowed string, words []string) int {
	res := make(map[byte]bool)
	count := 0

	for i := 0; i < len(allowed); i++ {
		res[allowed[i]] = true
	}

	for i := 0; i < len(words); i++ {
		found := false
		for j := 0; j < len(words[i]); j++ {
			if _, ok := res[words[i][j]]; !ok {
				found = true
                break
			}
		}

		if !found {
			count++
		}
	}

	return count
}