package leetcode

func closeStrings(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}

	freq1 := make(map[rune]int)
	freq2 := make(map[rune]int)

	for _, ch := range word1 {
		freq1[ch]++
	}

	for _, ch := range word2 {
		freq2[ch]++
	}

	if len(freq1) != len(freq2) {
		return false
	}

	for ch := range freq1 {
		if freq2[ch] == 0 {
			return false
		}
	}

	values1 := make([]int, 0, len(freq1))
	values2 := make([]int, 0, len(freq2))

	for _, value := range freq1 {
		values1 = append(values1, value)
	}

	for _, value := range freq2 {
		values2 = append(values2, value)
	}

	sort.Ints(values1)
	sort.Ints(values2)

	for i, value := range values1 {
		if values2[i] != value {
			return false
		}
	}

	return true
}
