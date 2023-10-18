package leetcode

import "strings"

func mostWordsFound(sentences []string) int {
	max := 0
	for i := 0; i < len(sentences); i++ {
		words := len(strings.Fields(sentences[i]))
		if max < words {
			max = words
		}
	}
	return max
}
