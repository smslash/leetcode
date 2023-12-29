package leetcode

func reverseWords(s string) string {
	rev := strings.Fields(s)
	res := make([]string, len(rev))

	for i := 0; i < len(rev); i++ {
		res[i] = rev[len(rev)-i-1]
	}

	return strings.Join(res, " ")
}
