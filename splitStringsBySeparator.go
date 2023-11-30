package leetcode

func splitWordsBySeparator(words []string, separator byte) []string {
	res := ""
    for i := 0; i < len(words); i++ {
		res += strings.ReplaceAll(words[i], string(separator), " ") + " "
	}
	return strings.Fields(res)
}
