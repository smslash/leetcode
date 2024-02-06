package leetcode

func findWords(words []string) []string {
	count1 := make([]int, len(words))
	count2 := make([]int, len(words))
	count3 := make([]int, len(words))

	for i := 0; i < len(words); i++ {
		for j := 0; j < len(words[i]); j++ {
			if strings.Contains("qwertyuiopQWERTYUIOP", string(words[i][j])) {
				count1[i]++
			} else if strings.Contains("asdfghjklASDFGHJKL", string(words[i][j])) {
				count2[i]++
			} else if strings.Contains("zxcvbnmZXCVBNM", string(words[i][j])) {
				count3[i]++
			}
		}
	}

	res := make([]string, 0)

	for i := 0; i < len(words); i++ {
		if count1[i] == len(words[i]) {
			res = append(res, words[i])
		} else if count2[i] == len(words[i]) {
			res = append(res, words[i])
		} else if count3[i] == len(words[i]) {
			res = append(res, words[i])
		}
	}

	return res
}
