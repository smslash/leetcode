package leetcode

func makeEqual(words []string) bool {
	res := make(map[byte]int)

	for i := 0; i < len(words); i++ {
		for j := 0; j < len(words[i]); j++ {
			res[words[i][j]]++
		}
	}

	for _, v := range res {
		if v%len(words) != 0 {
			return false
		}
	}

	return true
}
