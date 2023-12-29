package leetcode

func mergeAlternately(word1 string, word2 string) string {
	var min int
	if len(word1) > len(word2) {
		min = len(word2)
	} else {
		min = len(word1)
	}

	var arr []byte
	for i := 0; i < min; i++ {
		arr = append(arr, word1[i])
		arr = append(arr, word2[i])
	}

	if len(word1) != min {
		return string(arr) + word1[min:]
	}

	return string(arr) + word2[min:]
}
