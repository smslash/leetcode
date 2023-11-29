package leetcode

func hammingWeight(num uint32) int {
	s := strconv.FormatUint(uint64(num), 2)
	count := 0

	for i := 0; i < len(s); i++ {
		if s[i] == '1' {
			count++
		}
	}
	
	return count
}
