package leetcode

func minFlips(target string) int {
	res := 0
	for i := 0; i < len(target); i++ {
		s := target[i]
		if s == '1' {
			for i+1 < len(target) && target[i+1] == s {
				i++
			}
			res += 2
		}
	}
	if target[len(target)-1] == '1' {
		res--
	}
	return res
}
