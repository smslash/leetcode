package leetcode

func lengthOfLongestSubstring(s string) int {
	hashMap := make(map[byte]int, 26)
	maxLen := 0
	start := 0

	for i := 0; i < len(s); i++ {
		val := s[i] - 'a'
		if hashMap[val] > start {
			start = hashMap[val]
		}
		seqLen := i - start + 1
		if seqLen > maxLen {
			maxLen = seqLen
		}
		hashMap[val] = i + 1
	}
	return maxLen
}
