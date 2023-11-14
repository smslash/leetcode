package leetcode

func countPalindromicSubsequence(s string) int {
	posMap := make(map[rune][]int)
	for i, c := range s {
		posMap[c] = append(posMap[c], i)
	}

	count := 0
	for _, positions := range posMap {
		if len(positions) < 2 {
			continue
		}
		first, last := positions[0], positions[len(positions)-1]
		uniqueChars := make(map[rune]bool)
		for _, c := range s[first+1 : last] {
			uniqueChars[c] = true
		}
		count += len(uniqueChars)
	}
	return count
}
