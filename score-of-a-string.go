package leetcode

func scoreOfString(s string) int {
	var res int
	for i := 0; i < len(s)-1; i++ {
		if s[i] > s[i+1] {
			res += int(s[i]) - int(s[i+1])
		} else {
			res += int(s[i+1]) - int(s[i])
		}
	}
	return res
}
